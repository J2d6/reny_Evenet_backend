package reposiory

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
    
    tarifsJSON, err := json.Marshal(req.Tarifs)
    if err != nil {
        return uuid.Nil, err
    }
    
    // Préparation des fichiers pour JSON
    var fichiersJSON []byte
    if len(req.Fichiers) > 0 {
        fichiersPourJSON := make([]map[string]interface{}, len(req.Fichiers))
        for i, fichier := range req.Fichiers {
            // Décoder le base64 en bytes pour PostgreSQL
            donneesBytes, err := base64.StdEncoding.DecodeString(fichier.DonneesBase64)
            if err != nil {
                return uuid.Nil, err
            }
            
            fichiersPourJSON[i] = map[string]interface{}{
                "nom_fichier":  fichier.NomFichier,
                "type_mime":    fichier.TypeMime,
                "type_fichier": fichier.TypeFichier,
                "donnees_binaire": donneesBytes,
            }
        }
        fichiersJSON, err = json.Marshal(fichiersPourJSON)
        if err != nil {
            return uuid.Nil, err
        }
    } else {
        fichiersJSON = []byte("[]")
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

func  (r *EvenementRepository) CloseConn()  {
	r.conn.Close(context.Background())
}


func NewEvenementRepository(conn *pgx.Conn) *EvenementRepository  {
	return &EvenementRepository{*conn}
}