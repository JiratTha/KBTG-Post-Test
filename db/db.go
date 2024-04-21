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