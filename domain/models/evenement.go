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
    
    // Données lieu
    Lieu        LieuInput    `json:"lieu"`
    
    // Tarifs
    Tarifs      []TarifInput `json:"tarifs"`
    
    // Fichiers (nouveau!)
    Fichiers    []FichierInput `json:"fichiers,omitempty"`
}

type FichierInput struct {
    NomFichier  string `json:"nom_fichier"`
    TypeMime    string `json:"type_mime"`
    TypeFichier string `json:"type_fichier"` // 'photo', 'affiche', 'document'
    DonneesBase64 string `json:"donnees_base64"` // Fichier encodé en base64
}