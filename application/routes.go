// Dans application/routes.go
package application

import (
	"github.com/J2d6/reny_event/domain/service"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(r chi.Router, evenementService *service.EvenementService, authService *service.AuthentificationService, jwtSecret string) {
	evenementHandler := NewEvenementHandler(evenementService)
	authHandler := NewAuthHandler(authService, jwtSecret)

	r.Route("/v1", func(r chi.Router) {
		// Routes publiques
		r.Route("/auth", func(r chi.Router) {
			r.Post("/connexion", authHandler.ConnexionHandler)
			r.Post("/deconnexion", authHandler.DeconnexionHandler)
		})

		// Routes événements publiques (lecture seule)
		r.Get("/evenements/{id}", evenementHandler.GetEvenementHandler)
		r.Get("/evenements/{evenementId}/fichiers/{fichierId}/contenu", evenementHandler.GetFichierContenuHandler)

		// Routes protégées (création seulement)
		r.Route("/", func(r chi.Router) {
			r.Use(authHandler.MiddlewareAuth) // JWT required
			r.Post("/evenements", evenementHandler.CreateEvenementHandler)
		})
	})
}