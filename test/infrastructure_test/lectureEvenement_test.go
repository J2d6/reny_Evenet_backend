package infrastructure_test

import (
	"errors"
	"testing"
	"github.com/J2d6/reny_event/domain/interfaces"
	domain_error "github.com/J2d6/reny_event/domain/errors"
	"github.com/J2d6/reny_event/infrastructure/db"
	"github.com/J2d6/reny_event/infrastructure/repository"
	"github.com/google/uuid"
)



func TestLectureEvenement(t *testing.T) { 
	t.Run("NOT FOUND EVENEMENT", func (t *testing.T) {
		evenement_id := uuid.New()
		repo := createRepository(t)
		_ , err := repo.GetEvenementByID(evenement_id)
		if err == nil {
			t.Errorf("Didn't get the SQL error : %v", err)
		}
	})

	t.Run("Get evenement by known ID", func (t *testing.T) {
		evenement_id := uuid.MustParse("fc142deb-73c7-4dbb-8f51-fe05a8231836")
		repo:= createRepository(t)
		_ , err := repo.GetEvenementByID(evenement_id)
		assertError(t, err)
	})
	
}

func assertError(t testing.TB, err error)  {
	t.Helper()
	if err != nil {
		t.Errorf("ERROR : %v", err)
	}
}

func assertSQLError(t testing.TB, err error)  {
	t.Helper()
	if !errors.Is(err, &domain_error.ErreurSQL{}) {
		t.Errorf("Didn't get SQL ERROR : %v", err)
	}
}


func createRepository(t testing.TB) interfaces.EvenementRepository  {
	t.Helper()
	conn, err := db.CreateNewPgxConnexion()
	if err != nil {
		t.Fatalf("Failed to connect to the database")
	}
	repo := repository.NewEvenementRepository(conn)
	return repo
}