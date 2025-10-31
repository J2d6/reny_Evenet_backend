package main

import (
	"context"
	"log"
	"net/http"
	"github.com/J2d6/reny_event/application"
	"github.com/J2d6/reny_event/domain/service"
	"github.com/J2d6/reny_event/infrastructure/db"
	"github.com/J2d6/reny_event/infrastructure/repository"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// Connexion base de données
	dbConn, err := db.CreateNewPgxConnexion()
	if err != nil {
		log.Fatal("Erreur connexion BD:", err)
	}
	defer dbConn.Close(context.Background())

	// Initialisation des repositories
	evenementRepo := repository.NewEvenementRepository(dbConn)
	utilisateurRepo := repository.NewUtilisateurRepository(dbConn)

	// Initialisation des services
	evenementService := service.NewEvenementService(evenementRepo)
	authService := service.NewAuthentificationService(utilisateurRepo)

	// Clé JWT (à mettre dans une variable d'environnement en production)
	jwtSecret := "reny_evenet"
	if jwtSecret == "" {
		jwtSecret = "votre-cle-secrete-pour-developpement"
		log.Println("⚠️  Attention: Utilisation d'une clé JWT par défaut - À changer en production")
	}

	// Configuration du router
	r := chi.NewRouter()

	// Middlewares globaux
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)

	// Setup des routes avec authentification
	application.SetupRoutes(r, evenementService, authService, jwtSecret)

	// Démarrage serveur
	port := ":3000"
	log.Printf("🚀 Serveur démarré sur http://localhost%s", port)
	log.Printf("")
	log.Printf("🔐 Endpoints d'authentification:")
	log.Printf("   POST http://localhost%s/v1/auth/connexion", port)
	log.Printf("   POST http://localhost%s/v1/auth/deconnexion", port)
	log.Printf("")
	log.Printf("📝 Endpoints événements:")
	log.Printf("   POST http://localhost%s/v1/evenements (🔒 PROTÉGÉ - JWT requis)", port)
	log.Printf("   GET  http://localhost%s/v1/evenements/{id} (🔓 PUBLIC)", port)
	log.Printf("   GET  http://localhost%s/v1/evenements/{id}/fichiers/{id}/contenu (🔓 PUBLIC)", port)

	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatalf("❌ Erreur serveur: %v", err)
	}
}