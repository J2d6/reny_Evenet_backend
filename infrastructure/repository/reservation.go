package repository

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/J2d6/reny_event/domain/models"
)


func (repo EvenementRepository) Reserver(req models.ReservationRequest) (string, error) {
	
    var reservationID string
    
    // Convertir PlacesDemandees en JSON
    placesJSON, err := json.Marshal(req.PlacesDemandees)
    if err != nil {
        return "", fmt.Errorf("erreur marshaling JSON: %w", err)
    }
    
    // Appeler la fonction PostgreSQL
    err = repo.conn.QueryRow(
        context.Background(),
        "SELECT reserver_places($1, $2, $3)",
        req.Email,
        req.EvenementID,
        placesJSON,
    ).Scan(&reservationID)
    
    if err != nil {
        return "", fmt.Errorf("erreur lors de la réservation: %w", err)
    }
    
    return reservationID, nil
}
