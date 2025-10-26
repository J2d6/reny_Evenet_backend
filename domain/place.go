package domain

import "github.com/google/uuid"

type EtatPlace string

const (
	EtatDisponible EtatPlace = "disponible"
	EtatReservee   EtatPlace = "reservee"
	EtatVendue     EtatPlace = "vendue"
)

type Place struct {
	ID          uuid.UUID `json:"id"`
	Numero      string    `json:"numero"`
	Etat        EtatPlace `json:"etat"`
	EvenementID uuid.UUID `json:"evenement_id"`
	TarifID     uuid.UUID `json:"tarif_id"`

	// Relations
	Tarif *Tarif `json:"tarif,omitempty"`
}
