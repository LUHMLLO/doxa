package app

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Database struct {
	db *sql.DB
}

func NewDatabase() (*Database, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(time.Minute * 3)

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to PlanetScale")

	return &Database{
		db: db,
	}, nil
}

func (s *Database) InitializeTables() {
	if _, err := s.db.Exec(`
		CREATE TABLE IF NOT EXISTS Clients (
			id INT PRIMARY KEY,
			name VARCHAR(250),
			email VARCHAR(250) UNIQUE,
			phone VARCHAR(250) UNIQUE,
			created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			modified TIMESTAMP NULL
		);
	`); err != nil {
		log.Println(err)
	}

	if _, err := s.db.Exec(`
		CREATE TABLE IF NOT EXISTS Users (
			id INT PRIMARY KEY,
			username VARCHAR(250),
			password VARCHAR(250),
			avatar VARCHAR(250),
			role VARCHAR(250),
			created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			modified TIMESTAMP NULL,
			accessed TIMESTAMP NULL,
			client_id INT UNIQUE,
			FOREIGN KEY (client_id) REFERENCES Clients(id)
		);
	`); err != nil {
		log.Println(err)
	}
}
