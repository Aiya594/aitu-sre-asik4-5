package configs

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

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

	var db *sql.DB

	for i := 0; i < 10; i++ {
		db, err := sql.Open("postgres", connStr)
		if err == nil {
			err = db.Ping()
			if err == nil {
				return db, nil
			}
		}

		fmt.Println("DB not ready, retrying...")
		time.Sleep(2 * time.Second)
	}

	log.Print("Connected to DB!!")

	return db, nil
}
