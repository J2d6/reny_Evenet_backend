// Dans application/routes.go
package application

import (
	"github.com/J2d6/reny_event/application/handler"
	"github.com/J2d6/reny_event/domain/interfaces"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)


func SetupRoutes(r chi.Router, evenementService interfaces.EvenementService) {
	
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           86400, 
	}))

	// Route pour la version v1 de l'API
	r.Route("/v1", func(r chi.Router) {
		// Routes événements
		// r.Get("/evenements/{id}", evenementHandler.GetEvenementHandler)
		r.Get("/evenements/{id}", handler.GetEvenementByIDHandler(evenementService))
		r.Post("/evenements", handler.CreationEvenementHandler(evenementService)) 
		r.Post("/reservations", handler.ReserverHandler(evenementService)) 
	})

	// Route de santé
	// r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.WriteHeader(http.StatusOK)
	// 	w.Write([]byte(`{"status": "ok"}`))
	// })
}