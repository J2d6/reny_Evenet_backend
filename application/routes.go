// Dans application/routes.go
package application

import (
	"github.com/J2d6/reny_event/application/handler"
	"github.com/J2d6/reny_event/domain/interfaces"
	"github.com/go-chi/chi/v5"
)


func SetupRoutes(r chi.Router, evenementService interfaces.EvenementService) {
	
	// // Créer le handler
	// evenementHandler := NewEvenementHandler(evenementService)

	// Route pour la version v1 de l'API
	r.Route("/v1", func(r chi.Router) {
		// Routes événements
		// r.Get("/evenements/{id}", evenementHandler.GetEvenementHandler)
		r.Get("/evenements/{id}", handler.GetEvenementByIDHandler(evenementService))
		r.Post("/evenements", handler.CreationEvenementHandler(evenementService)) 
	})

	// Route de santé
	// r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.WriteHeader(http.StatusOK)
	// 	w.Write([]byte(`{"status": "ok"}`))
	// })
}