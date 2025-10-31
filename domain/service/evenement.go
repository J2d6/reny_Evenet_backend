package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/J2d6/reny_event/domain/interfaces"
	"github.com/J2d6/reny_event/domain/models"
	"github.com/google/uuid"
)

type EvenementService struct {
    repo interfaces.EvenementRepository
}

func NewEvenementService(repo interfaces.EvenementRepository) *EvenementService {
    return &EvenementService{repo: repo}
}

// CreerEvenement crée un événement complet avec lieu, tarifs et fichiers
func (s *EvenementService) CreerEvenement(
    ctx context.Context,
    req models.CreationEvenementRequest,
) (uuid.UUID, error) {
    
    // Validation des données
    if err := s.validerDonneesCreation(req); err != nil {
        return uuid.Nil, err
    }
    
    // Appel du repository pour créer l'événement
    evenementID, err := s.repo.CreateNewEvenement(ctx, req)
    if err != nil {
        return uuid.Nil, fmt.Errorf("échec création événement: %w", err)
    }
    
    return evenementID, nil
}

// validerDonneesCreation valide les règles métier pour la création
func (s *EvenementService) validerDonneesCreation(req models.CreationEvenementRequest) error {
    // Validation du titre
    if req.Titre == "" {
        return &ErreurValidation{Champ: "titre", Message: "Le titre est obligatoire"}
    }
    if len(req.Titre) > 150 {
        return &ErreurValidation{Champ: "titre", Message: "Le titre ne peut pas dépasser 150 caractères"}
    }
    
    // Validation des dates
    if req.DateDebut.IsZero() {
        return &ErreurValidation{Champ: "date_debut", Message: "La date de début est obligatoire"}
    }
    if req.DateFin.IsZero() {
        return &ErreurValidation{Champ: "date_fin", Message: "La date de fin est obligatoire"}
    }
    if req.DateDebut.Before(time.Now()) {
        return &ErreurValidation{Champ: "date_debut", Message: "La date de début doit être dans le futur"}
    }
    if req.DateFin.Before(req.DateDebut) {
        return &ErreurValidation{Champ: "date_fin", Message: "La date de fin doit être après la date de début"}
    }
    
    // Validation du type
    if req.Type_evenement == "" {
        return &ErreurValidation{Champ: "type_evenement", Message: "Le type d'événement est obligatoire"}
    }
    
    // Validation du lieu
    if req.Lieu.Nom == "" {
        return &ErreurValidation{Champ: "lieu.nom", Message: "Le nom du lieu est obligatoire"}
    }
    if len(req.Lieu.Nom) > 150 {
        return &ErreurValidation{Champ: "lieu.nom", Message: "Le nom du lieu ne peut pas dépasser 150 caractères"}
    }
    if req.Lieu.Adresse == "" {
        return &ErreurValidation{Champ: "lieu.adresse", Message: "L'adresse du lieu est obligatoire"}
    }
    if req.Lieu.Ville == "" {
        return &ErreurValidation{Champ: "lieu.ville", Message: "La ville du lieu est obligatoire"}
    }
    if req.Lieu.Capacite <= 0 {
        return &ErreurValidation{Champ: "lieu.capacite", Message: "La capacité du lieu doit être positive"}
    }
    
    // Validation des tarifs
    if len(req.Tarifs) == 0 {
        return &ErreurValidation{Champ: "tarifs", Message: "Au moins un tarif doit être spécifié"}
    }
    
    totalPlaces := 0
    for i, tarif := range req.Tarifs {
        if tarif.TypePlaceID == uuid.Nil {
            return &ErreurValidation{
                Champ:   "tarifs",
                Message: fmt.Sprintf("Le type de place est obligatoire pour le tarif %d", i+1),
            }
        }
        if tarif.Prix < 0 {
            return &ErreurValidation{
                Champ:   "tarifs",
                Message: fmt.Sprintf("Le prix ne peut pas être négatif pour le tarif %d", i+1),
            }
        }
        if tarif.NombrePlaces <= 0 {
            return &ErreurValidation{
                Champ:   "tarifs",
                Message: fmt.Sprintf("Le nombre de places doit être positif pour le tarif %d", i+1),
            }
        }
        totalPlaces += tarif.NombrePlaces
    }
    
    // Vérification que le total des places ne dépasse pas la capacité
    if totalPlaces > req.Lieu.Capacite {
        return &ErreurValidation{
            Champ:   "tarifs",
            Message: fmt.Sprintf("Le total des places (%d) dépasse la capacité du lieu (%d)", totalPlaces, req.Lieu.Capacite),
        }
    }
    
    // Validation des fichiers (nouveau!)
    for i, fichier := range req.Fichiers {
        if fichier.NomFichier == "" {
            return &ErreurValidation{
                Champ:   "fichiers",
                Message: fmt.Sprintf("Le nom de fichier est obligatoire pour le fichier %d", i+1),
            }
        }
        if fichier.TypeFichier == "" {
            return &ErreurValidation{
                Champ:   "fichiers", 
                Message: fmt.Sprintf("Le type de fichier est obligatoire pour le fichier %d", i+1),
            }
        }
        if fichier.TypeFichier != "photo" && fichier.TypeFichier != "affiche" && fichier.TypeFichier != "document" {
            return &ErreurValidation{
                Champ:   "fichiers",
                Message: fmt.Sprintf("Type de fichier invalide pour le fichier %d: doit être 'photo', 'affiche' ou 'document'", i+1),
            }
        }
        if fichier.DonneesBase64 == "" {
            return &ErreurValidation{
                Champ:   "fichiers",
                Message: fmt.Sprintf("Les données du fichier sont obligatoires pour le fichier %d", i+1),
            }
        }
    }
    
    return nil
}

// ErreurValidation représente une erreur de validation métier
type ErreurValidation struct {
    Champ   string
    Message string
}

func (e *ErreurValidation) Error() string {
    return fmt.Sprintf("Erreur validation %s: %s", e.Champ, e.Message)
}




// GetEvenementDetail récupère les détails complets d'un événement
func (s *EvenementService) GetEvenementDetail(
    ctx context.Context,
    evenementID uuid.UUID,
) (*models.EvenementDetail, error) {
    
    if evenementID == uuid.Nil {
        return nil, &ErreurValidation{Champ: "id", Message: "ID d'événement invalide"}
    }
    
    
    row, err := s.repo.GetEvenementByID(ctx, evenementID)
    if err != nil {
        return nil, fmt.Errorf("erreur récupération événement: %w", err)
    }
    if row == nil {
        return nil, nil // Événement non trouvé
    }
    
    // Transformation des données brutes en modèle métier
    return s.transformRowToEvenementDetail(row)
}


func (s *EvenementService) transformRowToEvenementDetail(row *models.EvenementRow) (*models.EvenementDetail, error) {
    var detail models.EvenementDetail
    
    // Champs directs (plus besoin de conversion de dates!)
    detail.ID = row.EvenementID
    detail.Titre = row.Titre
    detail.Description = row.Description
    detail.DateDebut = row.DateDebut 
    detail.DateFin = row.DateFin     
    
    // Désérialisation JSON vers les modèles métier
    if err := json.Unmarshal(row.TypeEvenement, &detail.Type); err != nil {
        return nil, fmt.Errorf("erreur désérialisation type_evenement: %w", err)
    }
    
    if err := json.Unmarshal(row.Lieu, &detail.Lieu); err != nil {
        return nil, fmt.Errorf("erreur désérialisation lieu: %w", err)
    }
    
    if err := json.Unmarshal(row.Tarifs, &detail.Tarifs); err != nil {
        return nil, fmt.Errorf("erreur désérialisation tarifs: %w", err)
    }
    
    if err := json.Unmarshal(row.Fichiers, &detail.Fichiers); err != nil {
        return nil, fmt.Errorf("erreur désérialisation fichiers: %w", err)
    }
    
    if err := json.Unmarshal(row.Statistiques, &detail.Statistiques); err != nil {
        return nil, fmt.Errorf("erreur désérialisation statistiques: %w", err)
    }
    
    return &detail, nil
}




// GetFichierContenu récupère le contenu binaire d'un fichier
func (s *EvenementService) GetFichierContenu(
    ctx context.Context,
    evenementID uuid.UUID,
    fichierID uuid.UUID,
) (*models.FichierContenu, error) {
    
    // Validation des UUID
    if evenementID == uuid.Nil {
        return nil, &ErreurValidation{Champ: "evenement_id", Message: "ID d'événement invalide"}
    }
    if fichierID == uuid.Nil {
        return nil, &ErreurValidation{Champ: "fichier_id", Message: "ID de fichier invalide"}
    }
    
    // Appel du repository
    contenu, err := s.repo.GetFichierContenu(ctx, evenementID, fichierID)
    if err != nil {
        return nil, fmt.Errorf("erreur récupération contenu fichier: %w", err)
    }
    
    if contenu == nil {
        return nil, nil // Fichier non trouvé
    }
    
    return contenu, nil
}