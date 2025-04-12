package routes

import (
	"context"
	"go-backend/internal/app/api/handlers"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(r *chi.Mux) {
	registerPOST[handlers.CreateUserInput](
		context.Background(),
		r,
		handlers.GetCreateUserRoutes(),
		handlers.CreateUserHandler{},
	)
}
