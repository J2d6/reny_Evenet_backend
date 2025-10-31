package interfaces

import (
	"context"

	"github.com/J2d6/reny_event/domain/models"
)

type UtilisateurRepository interface {
	VerifierCredentials(ctx context.Context, login, motDePasse string) (*models.Utilisateur, error)
	GetUtilisateurByLogin(ctx context.Context, login string) (*models.Utilisateur, error)
	CreerUtilisateur(ctx context.Context, login, motDePasse string) (*models.Utilisateur, error)
	CloseConn()
}