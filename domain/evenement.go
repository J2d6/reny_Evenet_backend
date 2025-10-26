package domain

import (
	"time"
	"github.com/google/uuid"
)

type Evenement struct {
	ID          uuid.UUID     `json:"id"`
	Titre       string        `json:"titre"`
	Description string        `json:"description,omitempty"`
	DateDebut   time.Time     `json:"date_debut"`
	DateFin     time.Time     `json:"date_fin"`
	Heure       string        `json:"heure"`
	TypeID      uuid.UUID     `json:"type_id"`
	LieuID      uuid.UUID     `json:"lieu_id"`

	// Relations
	TypeEvenement *TypeEvenement `json:"type_evenement,omitempty"`
	Lieu          *Lieu          `json:"lieu,omitempty"`
	Tarifs        []Tarif        `json:"tarifs,omitempty"`
}
