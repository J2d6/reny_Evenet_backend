package domain

import "github.com/google/uuid"

type Tarif struct {
	ID             uuid.UUID `json:"id"`
	Prix           float64   `json:"prix"`
	NombrePlaces   int       `json:"nombre_places"`
	EvenementID    uuid.UUID `json:"evenement_id"`
	TypePlaceID    uuid.UUID `json:"type_place_id"`

	// Relations
	TypePlace *TypePlace `json:"type_place,omitempty"`
}
