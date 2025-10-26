package domain

import "github.com/google/uuid"

type TypeEvenement struct {
	ID          uuid.UUID `json:"id"`
	Nom         string    `json:"nom"`
	Description string    `json:"description,omitempty"`
}
