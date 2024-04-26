package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	_ "github.com/lib/pq" // Import your SQL driver
)

func executeSQLFile(db *sql.DB, filepath string) error {

	fileContent, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}

	sqlCommands := string(fileContent)
	_, err = db.Exec(sqlCommands)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	DatabaseUrl := os.Getenv("DatabaseUrl")
	if DatabaseUrl == "" {
		DatabaseUrl = "host=localhost port=5432 user=postgres password=postgres dbname=ktaxes sslmode=disable"
	}
	DB, _ := sql.Open("postgres", DatabaseUrl)
	if err := executeSQLFile(DB, "./allowance.sql"); err != nil {
		log.Fatalf("Failed to execute master_deduct.sql: %v", err)
	}
	if err := executeSQLFile(DB, "./personnel_deduction.sql"); err != nil {
		log.Fatalf("Failed to execute master_tax_level.sql: %v", err)
	}

	fmt.Println("Database initialized successfully!")
}
