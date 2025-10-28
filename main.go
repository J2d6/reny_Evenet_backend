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
)

func main() {
	// Connexion base de données
	dbConn := db.CreateNewPgxConnexion()
	defer dbConn.Close(context.Background())

	// Repository
	evenementRepo := repository.NewEvenementRepository(dbConn)

	// Service
	evenementService := service.NewEvenementService(evenementRepo)

	// Handler
	evenementHandler := application.NewEvenementHandler(evenementService)

	// Router Chi
	r := chi.NewRouter()

	// Routes
	r.Route("/v1", func(r chi.Router) {
		r.Post("/evenements", evenementHandler.CreateEvenementHandler)  // POST création
		r.Get("/evenements/{id}", evenementHandler.GetEvenementHandler) // GET lecture
	})

	// Démarrage serveur
	port := ":3000"
	log.Printf("🚀 Serveur démarré sur http://localhost%s", port)
	log.Printf("📝 Endpoints:")
	log.Printf("   POST http://localhost%s/v1/evenements", port)
	log.Printf("   GET  http://localhost%s/v1/evenements/{id}", port)

	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatalf("❌ Erreur serveur: %v", err)
	}
}