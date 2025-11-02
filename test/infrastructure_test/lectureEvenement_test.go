package infrastructure_test

import (
	"testing"

	"github.com/J2d6/reny_event/infrastructure/db"
	"github.com/J2d6/reny_event/infrastructure/repository"
	"github.com/google/uuid"
)



func TestLectureEvenement(t *testing.T) { 
	t.Run("NOT FOUND EVENEMENT", func (t *testing.T) {
		evenement_id := uuid.New()
		conn, err := db.CreateNewPgxConnexion()
		if err != nil {
			t.Fatalf("Failed to connect to the database")
		}
		repo := repository.NewEvenementRepository(conn)

		_ , err = repo.GetEvenementByID(evenement_id)
		if err == nil {
			t.Errorf("ERROR : %v", err)
		}
	})

	t.Run("Get evenement by known ID", func (t *testing.T) {
		evenement_id := uuid.MustParse("fc142deb-73c7-4dbb-8f51-fe05a8231836")
		conn, err := db.CreateNewPgxConnexion()
		if err != nil {
			t.Fatalf("Failed to connect to the database")
		}
		repo := repository.NewEvenementRepository(conn)

		_ , err = repo.GetEvenementByID(evenement_id)
		if err != nil {
			t.Errorf("ERROR : %v", err)
		}
	})
	
}