package db

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func Connect() {
	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		// fallback to a local docker compose instance
		dbUrl = "postgres://postgres:password@localhost:5432/report_db?sslmode=disable"
	}

	var err error
	DB, err = sqlx.Connect("postgres", dbUrl)
	if err != nil {
		log.Println("Database connection failed, will use mock if needed:", err)
		return
	}
	log.Println("Connected to PostgreSQL")

	err = Migrate()
	if err != nil {
		log.Println("Migration failed:", err)
	}
}
