package repository

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/J2d6/reny_event/domain/models"
	"github.com/google/uuid"
)

func (repo EvenementRepository) GetEvenementByID(id uuid.UUID) (*models.EvenementComplet, error) {
	query := `SELECT obtenir_evenement_par_id($1) as evenement_data`
	var jsonData []byte
	err := repo.conn.QueryRow(context.Background(), query, id).Scan(&jsonData)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la récupération de l'événement: %w", err)
	}


	// Vérifier si c'est une erreur
	var errorResponse struct {
		Erreur bool   `json:"erreur"`
		Message string `json:"message"`
	}
	
	if err := json.Unmarshal(jsonData, &errorResponse); err == nil && errorResponse.Erreur {
		return nil, fmt.Errorf("erreur SQL: %s", errorResponse.Message)
	}

	var evenement models.EvenementComplet
	if err := json.Unmarshal(jsonData, &evenement); err != nil {
		return nil, fmt.Errorf("erreur de décodage JSON: %w", err)
	}

	
	return &evenement, nil
}