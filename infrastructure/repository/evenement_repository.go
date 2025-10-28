package repository

import (
	"context"
	"encoding/base64"
	"encoding/json"

	Int "github.com/J2d6/reny_event/domain/interfaces"
	"github.com/J2d6/reny_event/domain/models"
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
    
    tarifsJSON, err := r.prepareTarifsJSON(req.Tarifs)
    if err != nil {
        return uuid.Nil, err
    }
    
    fichiersJSON, err := r.prepareFichiersJSON(req.Fichiers)
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
        
        // Fichiers
        fichiersJSON,
    ).Scan(&nouvelID)
    
    if err != nil {
        return uuid.Nil, err
    }
    
    return nouvelID, nil
}


func (r *EvenementRepository) prepareTarifsJSON(tarifs []models.TarifInput) ([]byte, error) {
    return json.Marshal(tarifs)
}


func (r *EvenementRepository) prepareFichiersJSON(fichiers []models.FichierInput) ([]byte, error) {
    if len(fichiers) == 0 {
        return []byte("[]"), nil
    }
    
    fichiersPourJSON := make([]map[string]interface{}, len(fichiers))
    for i, fichier := range fichiers {
        // Décoder le base64 en bytes pour PostgreSQL
        donneesBytes, err := base64.StdEncoding.DecodeString(fichier.DonneesBase64)
        if err != nil {
            return nil, err
        }
        
        fichiersPourJSON[i] = map[string]interface{}{
            "nom_fichier":    fichier.NomFichier,
            "type_mime":      fichier.TypeMime,
            "type_fichier":   fichier.TypeFichier,
            "donnees_binaire": donneesBytes,
        }
    }
    
    return json.Marshal(fichiersPourJSON)
}

func (r *EvenementRepository) CloseConn() {
    r.conn.Close(context.Background())
}

func NewEvenementRepository(conn *pgx.Conn) *EvenementRepository {
    return &EvenementRepository{*conn}
}



// LECTURE 

func (r *EvenementRepository) GetEvenementByID(
    ctx context.Context, 
    evenementID uuid.UUID,
) (*models.EvenementRow, error) {
    
    var row models.EvenementRow
    
    query := `
        SELECT 
            evenement_id,
            titre,
            description_evenement,
            date_debut,
            date_fin,
            type_evenement,
            lieu,
            tarifs,
            fichiers,
            statistiques
        FROM vue_evenement_complet 
        WHERE evenement_id = $1
    `
    
    err := r.conn.QueryRow(ctx, query, evenementID).Scan(
        &row.EvenementID,
        &row.Titre,
        &row.Description,
        &row.DateDebut,     
        &row.DateFin,       
        &row.TypeEvenement,
        &row.Lieu,
        &row.Tarifs,
        &row.Fichiers,
        &row.Statistiques,
    )
    
    if err != nil {
        if err == pgx.ErrNoRows {
            return nil, nil // Événement non trouvé
        }
        return nil, err
    }
    
    return &row, nil
}