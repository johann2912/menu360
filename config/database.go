package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() *sql.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("❌ Error connecting to the database:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("❌ Database unreachable:", err)
	}

	DB = db
	log.Println("✅ Successful connection to PostgreSQL")
	return DB 
}
