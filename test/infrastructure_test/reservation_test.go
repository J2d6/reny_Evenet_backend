package infrastructure

import (
	// "encoding/json"
	// "fmt"
	"testing"

	"github.com/J2d6/reny_event/domain/interfaces"
	"github.com/J2d6/reny_event/domain/models"
	"github.com/J2d6/reny_event/infrastructure/db"
	"github.com/J2d6/reny_event/infrastructure/repository"
)



func TestReservation(t *testing.T) {
	repo := CreateRepository(t)


	req := models.ReservationRequest{
		Email:       "client@example.com",
		EvenementID: "fc142deb-73c7-4dbb-8f51-fe05a8231836",
		PlacesDemandees: []models.TypePlaceDemande{
			{
				TypePlaceID: "14467104-6d39-445c-a2d1-4dd1f697ac68",
				Nombre:      50,
			},
		},
	}

	// placesJSON, err := json.Marshal(req.PlacesDemandees)
    // if err != nil {
    //     return "", fmt.Errorf("erreur marshaling JSON: %w", err)
    // }


	reservation_id , err := repo.Reserver(req)
	if reservation_id == "" {
		t.Errorf("FAILED TO RESERVE: %v", err)
	}
}


func CreateRepository(t testing.TB) interfaces.EvenementRepository  {
	t.Helper()
	conn, err := db.CreateNewPgxConnexion()
	if err != nil {
		t.Fatalf("Failed to connect to the database")
	}
	repo := repository.NewEvenementRepository(conn)
	return repo
}