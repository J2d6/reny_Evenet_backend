package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)


func CreateNewPgxConnexion() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), 	"postgres://postgres:BoissonXXLenergy261001..@localhost:5432/postgres")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return conn
}