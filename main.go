package main

import (
	"context"
	"fmt"
	"github.com/JiratTha/assessment-tax/db"
	_ "github.com/JiratTha/assessment-tax/docs"
	"github.com/JiratTha/assessment-tax/router"
	"github.com/JiratTha/assessment-tax/util"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoswagger "github.com/swaggo/echo-swagger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	e := echo.New()

	DatabaseUrl := os.Getenv("DatabaseUrl")
	if DatabaseUrl == "" {
		DatabaseUrl = "host=localhost port=5432 user=postgres password=postgres dbname=ktaxes sslmode=disable"
		err := db.InitDB(DatabaseUrl)
		if err != nil {
			e.Logger.Fatal("Error fetching allowance:", err)
		}
	}

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080" // Default port if not specified
	}
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Secure())
	e.Use(middleware.Recover())
	e.Validator = &util.CustomValidator{Validator: validator.New()}
	e.GET("/swagger/*", echoswagger.WrapHandler)

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "healthy")
	})

	router.InitRoutes(e)

	go func() {
		if err := e.Start(":" + PORT); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server: ", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal("error shutting down the server: ", err)
	}
	e.Logger.Info("Server gracefully stopped")
	fmt.Println("shutting down the server")

}
