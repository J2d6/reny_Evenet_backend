package repository

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/J2d6/reny_event/domain/models"
	"github.com/google/uuid"
)


func (repo EvenementRepository) Reserver(req models.ReservationRequest) (string, error) {
	
    var reservationID string
    placesJSON, err := json.Marshal(req.PlacesDemandees)
    if err != nil {
        return "", fmt.Errorf("erreur marshaling JSON: %w", err)
    }
    
    err = repo.conn.QueryRow(
        context.Background(),
        RESERVER_QUERY,
        req.Email,
        req.EvenementID,
        placesJSON,
    ).Scan(&reservationID)
    
    if err != nil {
        return "", fmt.Errorf("erreur lors de la réservation: %w", err)
    }
    
    return reservationID, nil
}

func (repo EvenementRepository) GetAllReservationsFor(id_evenement uuid.UUID)  (string, error){
    var result string
    err := repo.conn.QueryRow(
        context.Background(),
        "SELECT obtenir_reservations_evenement($1)",
        id_evenement,
    ).Scan(&result)
    
    if err != nil {
        return "", err
    }
    
    return result,nil
}
