package service

import (
	"context"
	"fmt"
	"time"
	"github.com/J2d6/reny_event/domain/interfaces"
	"github.com/J2d6/reny_event/domain/models"
	"github.com/J2d6/reny_event/domain/errors"
	"github.com/google/uuid"
)

type EvenementService struct {
    repo interfaces.EvenementRepository
}

func NewEvenementService(repo interfaces.EvenementRepository) *EvenementService {
    return &EvenementService{repo: repo}
}


func (service *EvenementService) CreerEvenement(
    ctx context.Context,
    req models.CreationEvenementRequest,
) (uuid.UUID, error) {
    
    if err := service.validerDonneesCreation(req); err != nil {
        return uuid.Nil, err
    }
    
    evenementID, err := service.repo.CreateNewEvenement(context.Background(),req)
    if err != nil {
        return uuid.Nil, err
    }
    
    return evenementID, nil
}


func (service *EvenementService) validerDonneesCreation(req models.CreationEvenementRequest) error {
    // Validation du titre
    if req.Titre == "" {
        return &errors.ErreurValidation{Champ: "titre", Message: "Le titre est obligatoire"}
    }
    if len(req.Titre) > 150 {
        return &errors.ErreurValidation{Champ: "titre", Message: "Le titre ne peut pas dépasser 150 caractères"}
    }
    
    // Validation des dates
    if req.DateDebut.IsZero() {
        return &errors.ErreurValidation{Champ: "date_debut", Message: "La date de début est obligatoire"}
    }
    if req.DateFin.IsZero() {
        return &errors.ErreurValidation{Champ: "date_fin", Message: "La date de fin est obligatoire"}
    }
    if req.DateDebut.Before(time.Now()) {
        return &errors.ErreurValidation{Champ: "date_debut", Message: "La date de début doit être dans le futur"}
    }
    if req.DateFin.Before(req.DateDebut) {
        return &errors.ErreurValidation{Champ: "date_fin", Message: "La date de fin doit être après la date de début"}
    }
    
    // Validation du type
    if req.Type_evenement == "" {
        return &errors.ErreurValidation{Champ: "type_id", Message: "Le type d'événement est obligatoire"}
    }
    
    // Validation du lieu
    if req.Lieu.Nom == "" {
        return &errors.ErreurValidation{Champ: "lieu.nom", Message: "Le nom du lieu est obligatoire"}
    }
    if len(req.Lieu.Nom) > 150 {
        return &errors.ErreurValidation{Champ: "lieu.nom", Message: "Le nom du lieu ne peut pas dépasser 150 caractères"}
    }
    if req.Lieu.Adresse == "" {
        return &errors.ErreurValidation{Champ: "lieu.adresse", Message: "L'adresse du lieu est obligatoire"}
    }
    if req.Lieu.Ville == "" {
        return &errors.ErreurValidation{Champ: "lieu.ville", Message: "La ville du lieu est obligatoire"}
    }
    if req.Lieu.Capacite <= 0 {
        return &errors.ErreurValidation{Champ: "lieu.capacite", Message: "La capacité du lieu doit être positive"}
    }
    
    // Validation des tarifs
    if len(req.Tarifs) == 0 {
        return &errors.ErreurValidation{Champ: "tarifs", Message: "Au moins un tarif doit être spécifié"}
    }
    
    totalPlaces := 0
    for i, tarif := range req.Tarifs {
        if tarif.TypePlaceID == uuid.Nil {
            return &errors.ErreurValidation{
                Champ:   "tarifs",
                Message: fmt.Sprintf("Le type de place est obligatoire pour le tarif %d", i+1),
            }
        }
        if tarif.Prix < 0 {
            return &errors.ErreurValidation{
                Champ:   "tarifs",
                Message: fmt.Sprintf("Le prix ne peut pas être négatif pour le tarif %d", i+1),
            }
        }
        if tarif.NombrePlaces <= 0 {
            return &errors.ErreurValidation{
                Champ:   "tarifs",
                Message: fmt.Sprintf("Le nombre de places doit être positif pour le tarif %d", i+1),
            }
        }
        totalPlaces += tarif.NombrePlaces
    }
    
    // Vérification que le total des places ne dépasse pas la capacité
    if totalPlaces > req.Lieu.Capacite {
        return &errors.ErreurValidation{
            Champ:   "tarifs",
            Message: fmt.Sprintf("Le total des places (%d) dépasse la capacité du lieu (%d)", totalPlaces, req.Lieu.Capacite),
        }
    }
    
    return nil
}


