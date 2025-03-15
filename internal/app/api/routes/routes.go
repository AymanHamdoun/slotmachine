package routes

import (
	"context"
	"github.com/go-chi/chi/v5"
	"go-backend/internal/app/api/handlers"
)

func SetupRoutes(r *chi.Mux) {
	registerPOST[handlers.CreateUserInput](
		context.Background(),
		r,
		handlers.GetCreateUserRoutes(),
		&handlers.CreateUserHandler{},
	)
}
