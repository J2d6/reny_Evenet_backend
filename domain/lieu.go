package domain

import "github.com/google/uuid"

type Lieu struct {
	ID       uuid.UUID `json:"id"`
	Nom      string    `json:"nom"`
	Adresse  string    `json:"adresse"`
	Ville    string    `json:"ville"`
	Capacite int       `json:"capacite"`
}
