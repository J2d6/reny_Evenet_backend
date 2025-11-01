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
	"github.com/go-chi/cors"
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

	// Middleware CORS CONFIGURÉ POUR SERVEO
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*.serveo.net", "http://localhost:*", "https://localhost:*", "*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "X-Requested-With"},
		ExposedHeaders:   []string{"Link", "Content-Disposition"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Middlewares globaux
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)

	// Setup des routes avec authentification
	application.SetupRoutes(r, evenementService, authService, jwtSecret)

	// Route health check pour tester
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "OK", "message": "API is running with CORS enabled"}`))
	})

	// Démarrage serveur
	port := ":3000"
	log.Printf("🚀 Serveur démarré sur http://localhost%s", port)
	log.Printf("🌍 CORS configuré pour : *.serveo.net, localhost, et toutes les origines")
	log.Printf("")
	log.Printf("🔐 Endpoints d'authentification:")
	log.Printf("   POST http://localhost%s/v1/auth/connexion", port)
	log.Printf("   POST http://localhost%s/v1/auth/deconnexion", port)
	log.Printf("")
	log.Printf("📝 Endpoints événements:")
	log.Printf("   POST http://localhost%s/v1/evenements (🔒 PROTÉGÉ - JWT requis)", port)
	log.Printf("   GET  http://localhost%s/v1/evenements/{id} (🔓 PUBLIC)", port)
	log.Printf("   GET  http://localhost%s/v1/evenements/{id}/fichiers/{id}/contenu (🔓 PUBLIC)", port)
	log.Printf("")
	log.Printf("❤️  Health Check:")
	log.Printf("   GET  http://localhost%s/health", port)

	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatalf("❌ Erreur serveur: %v", err)
	}
}