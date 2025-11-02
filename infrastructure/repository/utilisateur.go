package repository

// import (
// 	"context"
// 	"fmt"

// 	"github.com/J2d6/reny_event/domain/models"
	
// 	"github.com/jackc/pgx/v5"
// )

// type UtilisateurRepository struct {
// 	conn *pgx.Conn
// }

// func NewUtilisateurRepository(conn *pgx.Conn) *UtilisateurRepository {
// 	return &UtilisateurRepository{conn: conn}
// }

// // VerifierCredentials vérifie le login/mot de passe et retourne l'utilisateur si valide
// func (r *UtilisateurRepository) VerifierCredentials(
//     ctx context.Context, 
//     login, motDePasse string,
// ) (*models.Utilisateur, error) {
    
//     var utilisateur models.Utilisateur
    
//     query := `
//         SELECT id, login, mot_de_passe
//         FROM utilisateur 
//         WHERE login = $1 AND mot_de_passe = $2
//     `
    
//     err := r.conn.QueryRow(ctx, query, login, motDePasse).Scan(
//         &utilisateur.ID,
//         &utilisateur.Login,
//         &utilisateur.MotDePasse,
//     )
    
//     if err != nil {
//         if err == pgx.ErrNoRows {
//             return nil, nil // Credentials invalides
//         }
//         return nil, fmt.Errorf("erreur vérification credentials: %w", err)
//     }
    
//     return &utilisateur, nil
// }

// // GetUtilisateurByLogin récupère un utilisateur par son login (pour vérification existence)
// func (r *UtilisateurRepository) GetUtilisateurByLogin(
//     ctx context.Context, 
//     login string,
// ) (*models.Utilisateur, error) {
    
//     var utilisateur models.Utilisateur
    
//     query := `
//         SELECT id, login, mot_de_passe
//         FROM utilisateur 
//         WHERE login = $1
//     `
    
//     err := r.conn.QueryRow(ctx, query, login).Scan(
//         &utilisateur.ID,
//         &utilisateur.Login,
//         &utilisateur.MotDePasse,
//     )
    
//     if err != nil {
//         if err == pgx.ErrNoRows {
//             return nil, nil // Utilisateur non trouvé
//         }
//         return nil, fmt.Errorf("erreur récupération utilisateur: %w", err)
//     }
    
//     return &utilisateur, nil
// }

// // CreerUtilisateur crée un nouvel utilisateur
// func (r *UtilisateurRepository) CreerUtilisateur(
//     ctx context.Context, 
//     login, motDePasse string,
// ) (*models.Utilisateur, error) {
    
//     var utilisateur models.Utilisateur
    
//     query := `
//         INSERT INTO utilisateur (login, mot_de_passe)
//         VALUES ($1, $2)
//         RETURNING id, login, mot_de_passe
//     `
    
//     err := r.conn.QueryRow(ctx, query, login, motDePasse).Scan(
//         &utilisateur.ID,
//         &utilisateur.Login,
//         &utilisateur.MotDePasse,
//     )
    
//     if err != nil {
//         return nil, fmt.Errorf("erreur création utilisateur: %w", err)
//     }
    
//     return &utilisateur, nil
// }

// func (r *UtilisateurRepository) CloseConn() {
//     r.conn.Close(context.Background())
// }