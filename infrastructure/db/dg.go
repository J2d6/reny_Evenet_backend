package db

import (
	"context"
	"github.com/jackc/pgx/v5"
)


func CreateNewPgxConnexion() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), 	"postgresql://postgres:gF7dYGWDK9tOUzCN@db.bbfsckuzadzzsdymgzmj.supabase.co:5432/postgres")
	if err != nil {
		return nil, err
	}

	return conn, nil
}