package main

import (
	"context"
	"log"
	"net/http"

	"github.com/J2d6/reny_event/domain/service"
	"github.com/J2d6/reny_event/infrastructure/db"
	"github.com/J2d6/reny_event/application"
	"github.com/J2d6/reny_event/infrastructure/repository"
)

func main() {

	dbConn := db.CreateNewPgxConnexion()
	defer dbConn.Close(context.Background())

	
	evenementRepo := reposiory.NewEvenementRepository(dbConn)

	
	evenementService := service.NewEvenementService(evenementRepo)

	
	evenementHandler := application.NewEvenementHandler(evenementService)


	http.HandleFunc("/v1/evenements", evenementHandler.CreateEvenementHandler)

	
	port := ":3000"
	log.Printf("🚀 Serveur démarré sur http://localhost%s", port)
	log.Printf("📝 Endpoint: POST http://localhost%s/v1/evenements", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("❌ Erreur serveur: %v", err)
	}
}