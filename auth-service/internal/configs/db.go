package configs

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func NewDB() (*sql.DB, error) {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	pswd := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	mode := os.Getenv("DB_MODE")
	port := os.Getenv("DB_PORT")

	connStr := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host, user, pswd, name, port, mode,
	)

	log.Print("Connecting to DB..")

	return sql.Open("postgres", connStr)
}
