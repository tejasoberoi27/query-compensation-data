package main

import (
	"log"
	"query-compensation-data/internal/compensation/config"
	compensation_db "query-compensation-data/internal/compensation/db"
	"query-compensation-data/internal/compensation/env/postgres"
	"query-compensation-data/internal/compensation/helper/cleaner"
	"query-compensation-data/internal/compensation/helper/db_population"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	pg, err := postgres.New(cfg.Postgres)
	if err != nil {
		log.Fatalf("Failed to initialize PostgreSQL connection: %v", err)
	}

	h, err := pg.DB.GetHandler()
	if err != nil {
		log.Fatalf("Failed to get database handler: %v", err)
	}

	err = compensation_db.InitTable(h) // Initialize the database table
	if err != nil {
		log.Fatalf("Failed to initialize the database table: %v", err)
	}

	entries, err := cleaner.CleanAndParseCSV("assets/salary_survey-3.csv")
	if err != nil {
		log.Fatalf("Failed to clean and parse CSV: %v", err)
	}

	err = db_population.InsertData(h, entries)
	if err != nil {
		log.Fatalf("Failed to insert data into the database: %v", err)
	}

	log.Println("Data inserted successfully.")
}
