package main

import (
	"log"
	"net/http"

	"github.com/J2d6/reny_event/application"
	"github.com/J2d6/reny_event/domain/service"
	"github.com/J2d6/reny_event/infrastructure/db"
	"github.com/J2d6/reny_event/infrastructure/repository"
	"github.com/go-chi/chi/v5"
)

func main() {
	// Initialiser le routeur Chi
	r := chi.NewRouter()

	// Initialiser les services (pour l'instant vides)
	conn, err := db.CreateNewPgxConnexion()
	if err != nil {
		return
	}
	evenementRepositorry := repository.NewEvenementRepository(conn)
	evenementService := service.NewEvenementService(evenementRepositorry)
	// authService := &service.AuthentificationService{}
	// jwtSecret := "votre-secret-jwt-temporaire"

	// Configurer les routes
	application.SetupRoutes(r, evenementService)

	// Démarrer le serveur
	port := ":3000"
	log.Printf("Serveur démarré sur http://localhost%s", port)
	log.Printf("Endpoint création événement: POST http://localhost%s/v1/evenements", port)
	
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatalf("Erreur du serveur: %v", err)
	}
}




