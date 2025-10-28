package models

import (
    "time"
)


type CreationEvenementRequest struct {
    // Données événement
    Type_evenement string    `json:"type_id"`
    Titre       string       `json:"titre"`
    Description string       `json:"description"`
    DateDebut   time.Time    `json:"date_debut"`
    DateFin     time.Time    `json:"date_fin"`
    
    // Données lieu (plus de LieuID)
    Lieu        LieuInput    `json:"lieu"`
    
    // Tarifs
    Tarifs      []TarifInput `json:"tarifs"`
}