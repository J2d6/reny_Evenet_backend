package application

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/J2d6/reny_event/application/request"
	Int "github.com/J2d6/reny_event/domain/interfaces"
	"github.com/J2d6/reny_event/domain/models"
	"github.com/J2d6/reny_event/domain/service"
	"github.com/google/uuid"
)

type EvenementHandler struct {
    service *service.EvenementService
}

func NewEvenementHandler(service *service.EvenementService) *EvenementHandler {
    return &EvenementHandler{service: service}
}

func (h *EvenementHandler) CreateEvenementHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println("IN REQUEST")
    var apiReq request.CreateEvenementRequest
    
    // Décodage - Go convertit automatiquement les dates
    if err := json.NewDecoder(r.Body).Decode(&apiReq); err != nil {
        h.sendError(w, http.StatusBadRequest, "Format JSON invalide")
        return
    }
    
    // Transformation directe - pas besoin de parser les dates
    internalReq := h.transformToInternalModel(apiReq)
    
    // Appel du service
    evenementID, err := h.service.CreerEvenement(r.Context(), internalReq)
    if err != nil {
        h.sendError(w, http.StatusBadRequest, err.Error())
        return
    }
    
    // Réponse
    response := map[string]interface{}{
        "id": evenementID,
        "message": "Événement créé avec succès",
    }
    
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(response)
}

// transformToInternalModel - transformation simple
func (h *EvenementHandler) transformToInternalModel(apiReq request.CreateEvenementRequest) models.CreationEvenementRequest {
    // Transformation des tarifs
    tarifs := make([]models.TarifInput, len(apiReq.Tarifs))
    for i, tarifReq := range apiReq.Tarifs {
        typePlaceID := Int.TypePlaceIDMap[tarifReq.TypePlace]
        
        tarifs[i] = models.TarifInput{
            TypePlaceID:  uuid.MustParse(typePlaceID),
            Prix:         tarifReq.Prix,
            NombrePlaces: tarifReq.NombrePlaces,
        }
    }
    
    // Transformation des fichiers
    fichiers := make([]models.FichierInput, len(apiReq.Fichiers))
    for i, fichierReq := range apiReq.Fichiers {
        fichiers[i] = models.FichierInput{
            NomFichier:    fichierReq.NomFichier,
            TypeMime:      fichierReq.TypeMime,
            TypeFichier:   fichierReq.TypeFichier,
            DonneesBase64: fichierReq.DonneesBase64,
        }
    }
    
    return models.CreationEvenementRequest{
        Type_evenement: apiReq.TypeEvenement,
        Titre:          apiReq.Titre,
        Description:    apiReq.Description,
        DateDebut:      apiReq.DateDebut,      
        DateFin:        apiReq.DateFin,        
        Lieu: models.LieuInput{
            Nom:      apiReq.Lieu.Nom,
            Adresse:  apiReq.Lieu.Adresse,
            Ville:    apiReq.Lieu.Ville,
            Capacite: apiReq.Lieu.Capacite,
        },
        Tarifs:   tarifs,
        Fichiers: fichiers,
    }
}

func (h *EvenementHandler) sendError(w http.ResponseWriter, status int, message string) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    json.NewEncoder(w).Encode(map[string]string{
        "error": message,
    })
}