package infrastructure_test

import (
	"context"
	"encoding/base64"
	"os"
	"path/filepath"
	"testing"
	"time"

	repo "github.com/J2d6/reny_event/domain/interfaces"
	"github.com/J2d6/reny_event/domain/models"
	db "github.com/J2d6/reny_event/infrastructure/db"
	reposiory "github.com/J2d6/reny_event/infrastructure/repository"
	"github.com/google/uuid"
)


func TestCreateNewEvenementAvecVraisFichiers(t *testing.T) {

	ev_repo := reposiory.NewEvenementRepository(db.CreateNewPgxConnexion())
	
	// Chargement des fichiers réels
	fichiers, err := chargerFichiersDeTest(t)
	if err != nil {
		t.Fatalf("Erreur chargement fichiers: %v", err)
	}

	req := models.CreationEvenementRequest{
		Type_evenement: "Concert",
		Titre:          "Concert avec mage4",
		Description:    "Concert Mage4",
		DateDebut:      time.Date(2025, 11, 31, 20, 0, 0, 0, time.UTC),
		DateFin:        time.Date(2025, 11, 31, 23, 0, 0, 0, time.UTC),

		Lieu: models.LieuInput{
			Nom:      "stade Barea Mahamasia",
			Adresse:  "Mahamasina, Stade Barea",
			Ville:    "Antananarivo",
			Capacite: 3000,
		},

		Tarifs: []models.TarifInput{
			{
				TypePlaceID:  uuid.MustParse(repo.TypePlaceIDMap["VIP"]),
				Prix:         120.00,
				NombrePlaces: 50,
			},
			{
				TypePlaceID:  uuid.MustParse(repo.TypePlaceIDMap["Standard"]),
				Prix:         50.00,
				NombrePlaces: 150,
			},
			{
				TypePlaceID:  uuid.MustParse(repo.TypePlaceIDMap["Premium"]),
				Prix:         90.00,
				NombrePlaces: 90,
			},
		},

		Fichiers: fichiers,
	}

	evenementID, err := ev_repo.CreateNewEvenement(context.Background(), req)
	if err != nil {
		t.Errorf("Echec de creation de l'event : %v", err)
	}
	
	if evenementID == uuid.Nil {
		t.Error("L'ID de l'événement ne devrait pas être nil")
	} else {
		t.Logf("Événement créé avec succès - ID: %s", evenementID.String())
	}

	// Nettoyage
	defer ev_repo.CloseConn()
}

// chargerFichiersDeTest charge de vrais fichiers depuis le système de fichiers
func chargerFichiersDeTest(t testing.TB) ([]models.FichierInput, error) {
	t.Helper()
	// Crée un dossier test_files à la racine de ton projet
	dossierTest := "./test_files"
	
	var fichiers []models.FichierInput
	
	// Liste des fichiers à charger
	fichiersATester := []struct {
		chemin     string
		typeFichier string
	}{
		{"IMG_9997.JPG", "affiche"},
		{"IMG_9998.JPG", "photo"},
	}
	
	for _, f := range fichiersATester {
		cheminComplet := filepath.Join(dossierTest, f.chemin)
		
		// Vérifie si le fichier existe
		if _, err := os.Stat(cheminComplet); os.IsNotExist(err) {
			t.Logf("Fichier non trouvé: %s - création d'un fichier vide pour le test", cheminComplet)
			// Crée un fichier vide pour le test
			if err := creerFichierVide(cheminComplet); err != nil {
				return nil, err
			}
		}
		
		// Lit le fichier
		donnees, err := os.ReadFile(cheminComplet)
		if err != nil {
			return nil, err
		}
		
		// Détermine le type MIME basique
		typeMime := determinerTypeMIME(f.chemin)
		
		fichiers = append(fichiers, models.FichierInput{
			NomFichier:    filepath.Base(f.chemin),
			TypeMime:      typeMime,
			TypeFichier:   f.typeFichier,
			DonneesBase64: base64.StdEncoding.EncodeToString(donnees),
		})
		
		t.Logf("Fichier chargé: %s (%d bytes)", f.chemin, len(donnees))
	}
	
	return fichiers, nil
}

// creerFichierVide crée un petit fichier de test si il n'existe pas
func creerFichierVide(chemin string) error {
	// Crée le dossier parent si nécessaire
	dossier := filepath.Dir(chemin)
	if err := os.MkdirAll(dossier, 0755); err != nil {
		return err
	}
	
	// Contenu minimal selon l'extension
	var contenu []byte
	switch filepath.Ext(chemin) {
	case ".jpg", ".jpeg":
		// Header JPEG minimal
		contenu = []byte{0xFF, 0xD8, 0xFF, 0xE0}
	case ".png":
		// Header PNG minimal  
		contenu = []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}
	case ".pdf":
		// Header PDF minimal
		contenu = []byte("%PDF-1.4\n")
	default:
		contenu = []byte("contenu de test")
	}
	
	return os.WriteFile(chemin, contenu, 0644)
}

// determinerTypeMIME détermine le type MIME basé sur l'extension
func determinerTypeMIME(nomFichier string) string {
	ext := filepath.Ext(nomFichier)
	switch ext {
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".png":
		return "image/png"
	case ".pdf":
		return "application/pdf"
	case ".txt":
		return "text/plain"
	default:
		return "application/octet-stream"
	}
}