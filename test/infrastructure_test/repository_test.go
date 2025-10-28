package infrastructure_test

import (
	"context"
	"testing"
	"time"

	repo "github.com/J2d6/reny_event/domain/interfaces"
	"github.com/J2d6/reny_event/domain/models"
	db "github.com/J2d6/reny_event/infrastructure/db"
	reposiory "github.com/J2d6/reny_event/infrastructure/repository"
	"github.com/google/uuid"
)

func TestCreateNewEvenement(t *testing.T) {

	ev_repo := reposiory.NewEvenementRepository(db.CreateNewPgxConnexion())
	req := models.CreationEvenementRequest{
		Type_evenement : "Foire",
		Titre:       "Foire de La Grande Ile",
		Description: "foire nationale des artisans malgaches",
		DateDebut:   time.Date(2025, 11, 16, 20, 0, 0, 0, time.UTC),
		DateFin:     time.Date(2025, 11, 18, 23, 0, 0, 0, time.UTC),

		Lieu: models.LieuInput{
			Nom:      "stade Barea Mahamasia",
			Adresse:  "Mahamasina, Stade Barea",
			Ville:    "Antananarivo",
			Capacite: 3000,
		},

		Tarifs: []models.TarifInput{
			{
				TypePlaceID:  uuid.MustParse(repo.TypePlaceIDMap["VIP"]),
				Prix:         120.00,
				NombrePlaces: 50,
			},
			{
				TypePlaceID:  uuid.MustParse(repo.TypePlaceIDMap["Standard"]),
				Prix:         50.00,
				NombrePlaces: 150,
			},
			{
				TypePlaceID:  uuid.MustParse(repo.TypePlaceIDMap["Premium"]),
				Prix:         90.00,
				NombrePlaces: 90,
			},
		},
	}

	evenementID, err := ev_repo.CreateNewEvenement(context.Background(), req)
	if evenementID == uuid.Nil {
		t.Errorf("Echec de creation de l'event : %v", err)
	}

}
