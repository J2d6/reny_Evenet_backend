package db

import (
	"context"
	"github.com/jackc/pgx/v5"
)


func CreateNewPgxConnexion() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), 	"postgresql://reny_event:vepC19IWOqDgy68Va3p6tdtDO1WW1iYV@dpg-d44qlgbipnbc73aps8i0-a/reny_event")
	if err != nil {
		return nil, err
	}

	return conn, nil
}