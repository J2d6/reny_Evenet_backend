package infrastructure

import (
	"testing"
	"github.com/J2d6/reny_event/domain/models"
)



func TestReservation(t *testing.T) {
	t.Run("Reservation with success", func (t *testing.T) {
		repo := CreateRepository(t)

		req := models.ReservationRequest{
			Email:       "client@example.com",
			EvenementID: "fc142deb-73c7-4dbb-8f51-fe05a8231836",
			PlacesDemandees: []models.TypePlaceDemande{
				{
					TypePlaceID: "14467104-6d39-445c-a2d1-4dd1f697ac68",
					Nombre:      5,
				},
			},
		}
		reservation_id , err := repo.Reserver(req)
		if reservation_id == "" {
			t.Errorf("FAILED TO RESERVE: %v", err)
		}
	})

	t.Run("Reservation with no place dispo", func (t *testing.T) {
		repo := CreateRepository(t)

		req := models.ReservationRequest{
			Email:       "client@example.com",
			EvenementID: "fc142deb-73c7-4dbb-8f51-fe05a8231836",
			PlacesDemandees: []models.TypePlaceDemande{
				{
					TypePlaceID: "14467104-6d39-445c-a2d1-4dd1f697ac68",
					Nombre:      500,
				},
			},
		}
		_ , err := repo.Reserver(req)
		if err == nil {
			t.Errorf("FAILED TO RESERVE: %v", err)
		}
	})
		
}


