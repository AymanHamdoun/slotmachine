package routes

import (
	"go-backend/internal/app/api/handlers"
	"go-backend/internal/app/api/middleware"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(r *chi.Mux) {
	// User routes
	userHandler := handlers.NewUserHandler()

	r.Route("/api/v1", func(r chi.Router) {
		// Apply common middleware
		r.Use(middleware.BindInput)

		r.Route("/users", func(r chi.Router) {
			r.Post("/", userHandler.Create)
			// Add more user routes here
		})

		// Add more route groups here
	})
}
