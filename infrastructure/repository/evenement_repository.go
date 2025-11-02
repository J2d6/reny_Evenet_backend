package repository

import (
	"github.com/J2d6/reny_event/domain/interfaces"
	"github.com/jackc/pgx/v5"
)



type EvenementRepository struct {
	conn *pgx.Conn
}

func NewEvenementRepository(conn *pgx.Conn)  interfaces.EvenementRepository {
	return EvenementRepository{conn: conn}
}
