package db

import (
	"context"
	"github.com/jackc/pgx/v5"
)


func CreateNewPgxConnexion() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), 	"postgres://postgres:BoissonXXLenergy261001..@localhost:5432/reny_event")
	if err != nil {
		return nil, err
	}

	return conn, nil
}