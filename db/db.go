package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

var DB *sqlx.DB

func InitDB(dataSourceName string) error {
	var err error
	DB, err = sqlx.Connect("postgres", dataSourceName)
	if err != nil {
		log.Panic(err)
	}
	return nil
}

func SetDB(database *sqlx.DB) {
	DB = database
}
func CloseDB() {
	if DB != nil {
		err := DB.Close()
		if err != nil {
			log.Println("Failed to close database connection:", err)
		} else {
			log.Println("Database connection closed successfully.")
		}
	}
}
