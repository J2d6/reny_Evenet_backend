package models

type ReservationRequest struct {
    Email           string                   `json:"email"`
    EvenementID     string                   `json:"evenement_id"`
    PlacesDemandees []TypePlaceDemande       `json:"places_demandees"`
}

type TypePlaceDemande struct {
    TypePlaceID string `json:"type_place_id"`
    Nombre      int    `json:"nombre"`
}