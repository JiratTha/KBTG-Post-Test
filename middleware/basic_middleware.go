package middleware

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func loadEnv() error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}
	return nil
}

func BasicAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := loadEnv()
		if err != nil {
			return err
		}

		username := os.Getenv("ADMIN_USERNAME")
		password := os.Getenv("ADMIN_PASSWORD")

		if username == "" || password == "" {
			return errors.New("missing admin credentials")
		}

		if c.FormValue("username") == username && c.FormValue("password") == password {
			return next(c)
		}

		return echo.ErrUnauthorized
	}
}
