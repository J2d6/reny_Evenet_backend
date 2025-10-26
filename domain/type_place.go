package domain

import "github.com/google/uuid"

type TypePlace struct {
	ID         uuid.UUID `json:"id"`
	Nom        string    `json:"nom"`
	Description string   `json:"description,omitempty"`
	Avantages   string   `json:"avantages,omitempty"`
}
