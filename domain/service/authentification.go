// // Dans domain/service/authentification.go
package service

// import (
// 	"context"
// 	"fmt"

// 	"github.com/J2d6/reny_event/domain/interfaces"
// 	"github.com/J2d6/reny_event/domain/models"
// 	"github.com/J2d6/reny_event/domain/errors"
// )

// type AuthentificationService struct {
// 	utilisateurRepo interfaces.UtilisateurRepository
// }

// func NewAuthentificationService(utilisateurRepo interfaces.UtilisateurRepository) *AuthentificationService {
// 	return &AuthentificationService{
// 		utilisateurRepo: utilisateurRepo,
// 	}
// }

// // AuthentifierUtilisateur vérifie les credentials et retourne l'utilisateur si valide
// func (s *AuthentificationService) AuthentifierUtilisateur(
// 	ctx context.Context,
// 	login, motDePasse string,
// ) (*models.Utilisateur, error) {
	
// 	// Validation des paramètres
// 	if login == "" {
// 		return nil, &ErreurValidation{Champ: "login", Message: "Le login est obligatoire"}
// 	}
// 	if motDePasse == "" {
// 		return nil, &ErreurValidation{Champ: "mot_de_passe", Message: "Le mot de passe est obligatoire"}
// 	}

// 	// Vérification des credentials via le repository
// 	utilisateur, err := s.utilisateurRepo.VerifierCredentials(ctx, login, motDePasse)
// 	if err != nil {
// 		return nil, fmt.Errorf("erreur lors de l'authentification: %w", err)
// 	}

// 	if utilisateur == nil {
// 		return nil, &errors.ErreurAuthentification{Message: "Login ou mot de passe incorrect"}
// 	}

// 	return utilisateur, nil
// }

// // CreerUtilisateur crée un nouvel utilisateur
// func (s *AuthentificationService) CreerUtilisateur(
// 	ctx context.Context,
// 	login, motDePasse string,
// ) (*models.Utilisateur, error) {
	
// 	// Validation des paramètres
// 	if err := s.validerDonneesCreation(login, motDePasse); err != nil {
// 		return nil, err
// 	}

// 	// Vérifier si l'utilisateur existe déjà
// 	existant, err := s.utilisateurRepo.GetUtilisateurByLogin(ctx, login)
// 	if err != nil {
// 		return nil, fmt.Errorf("erreur vérification existence utilisateur: %w", err)
// 	}

// 	if existant != nil {
// 		return nil, &ErreurValidation{Champ: "login", Message: "Ce login est déjà utilisé"}
// 	}

// 	// Création de l'utilisateur
// 	utilisateur, err := s.utilisateurRepo.CreerUtilisateur(ctx, login, motDePasse)
// 	if err != nil {
// 		return nil, fmt.Errorf("erreur création utilisateur: %w", err)
// 	}

// 	return utilisateur, nil
// }

// // validerDonneesCreation valide les règles métier pour la création d'utilisateur
// func (s *AuthentificationService) validerDonneesCreation(login, motDePasse string) error {
// 	// Validation du login
// 	if login == "" {
// 		return &ErreurValidation{Champ: "login", Message: "Le login est obligatoire"}
// 	}
// 	if len(login) < 3 {
// 		return &ErreurValidation{Champ: "login", Message: "Le login doit contenir au moins 3 caractères"}
// 	}
// 	if len(login) > 50 {
// 		return &ErreurValidation{Champ: "login", Message: "Le login ne peut pas dépasser 50 caractères"}
// 	}

// 	// Validation du mot de passe
// 	if motDePasse == "" {
// 		return &ErreurValidation{Champ: "mot_de_passe", Message: "Le mot de passe est obligatoire"}
// 	}
// 	if len(motDePasse) < 4 {
// 		return &ErreurValidation{Champ: "mot_de_passe", Message: "Le mot de passe doit contenir au moins 4 caractères"}
// 	}

// 	return nil
// }

// // GetUtilisateurByLogin récupère un utilisateur par son login
// func (s *AuthentificationService) GetUtilisateurByLogin(
// 	ctx context.Context,
// 	login string,
// ) (*models.Utilisateur, error) {
	
// 	if login == "" {
// 		return nil, &ErreurValidation{Champ: "login", Message: "Le login est obligatoire"}
// 	}

// 	utilisateur, err := s.utilisateurRepo.GetUtilisateurByLogin(ctx, login)
// 	if err != nil {
// 		return nil, fmt.Errorf("erreur récupération utilisateur: %w", err)
// 	}

// 	return utilisateur, nil
// }