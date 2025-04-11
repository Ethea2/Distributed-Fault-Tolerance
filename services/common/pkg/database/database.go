package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

// DB_HOST=postgres
// DB_PORT=5432
// DB_USER=postgres
// DB_PASSWORD=postgrespassword
// DB_NAME=enrollment_system

func ConnectDB() (*pgx.Conn, error) {

	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	db_user := os.Getenv("DB_USER")
	db_password := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")
	db_schema := os.Getenv("DB_SCHEMA")

	db_uri := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?search_path=%s",
		db_user, db_password, db_host, db_port, db_name, db_schema)

	fmt.Printf("Connecting to database with URI: postgres://%s:***@%s:%s/%s?search_path=%s\n",
		db_user, db_host, db_port, db_name, db_schema)

	conn, err := pgx.Connect(context.Background(), db_uri)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to DB: %v\n", err)
		return nil, err
	}

	fmt.Println("Successfully connected to database")
	return conn, nil
}
