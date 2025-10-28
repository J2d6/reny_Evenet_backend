package reposiory

import (
	"context"
	"encoding/json"

	"github.com/J2d6/reny_event/domain/models"
	Int "github.com/J2d6/reny_event/domain/interfaces"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)


type EvenementRepository struct {
	conn pgx.Conn
}


func (r *EvenementRepository) CreateNewEvenement(
    ctx context.Context, 
    req models.CreationEvenementRequest,
) (uuid.UUID, error) {
    
    var nouvelID uuid.UUID
    
    tarifsJSON, err := json.Marshal(req.Tarifs)
    if err != nil {
        return uuid.Nil, err
    }
    
    
    err = r.conn.QueryRow(
        ctx,
        CREATE_EVENEMENT_COMPLET_QUERY,
        // Données événement
        req.Titre,
        req.Description,
        req.DateDebut,
        req.DateFin,
        Int.TypeEvenementIDMap[req.Type_evenement],
        
        // Données lieu
        req.Lieu.Nom,
        req.Lieu.Adresse,
        req.Lieu.Ville,
        req.Lieu.Capacite,
        
        // Tarifs
        tarifsJSON,
    ).Scan(&nouvelID)
    
    if err != nil {
        return uuid.Nil, err
    }
    
    return nouvelID, nil
}

func  (r *EvenementRepository) CloseConn()  {
	r.conn.Close(context.Background())
}


func NewEvenementRepository(conn *pgx.Conn) *EvenementRepository  {
	return &EvenementRepository{*conn}
}