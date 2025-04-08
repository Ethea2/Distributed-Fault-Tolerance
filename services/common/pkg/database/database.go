package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

// DB_HOST=postgres
// DB_PORT=5432
// DB_USER=postgres
// DB_PASSWORD=postgrespassword
// DB_NAME=enrollment_system

func ConnectDB() (*pgx.Conn, error) {
	godotenv.Load()

	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	db_user := os.Getenv("DB_USER")
	db_password := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")

	db_uri := db_host + "://" + db_user + ":" + db_password + "@localhost:" + db_port + "/" + db_name

	conn, err := pgx.Connect(context.Background(), db_uri)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to DB")
	}

	return conn, nil
}
