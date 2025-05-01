package routes

import (
	"context"
	"go-backend/internal/app/api/apihandlers"
	"go-backend/internal/app/api/viewhandlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(r *chi.Mux) {
	// Serves compiled / generated assets by npm build
	fs := http.FileServer(http.Dir("web/dist/assets"))
	r.Handle("/assets/*", http.StripPrefix("/assets/", fs))

	// Serves Manual assets like images
	publicFS := http.FileServer(http.Dir("web/public"))
	r.Handle("/public/*", http.StripPrefix("/public/", publicFS))

	registerPOST[apihandlers.CreateUserInput](
		context.Background(),
		r,
		apihandlers.GetCreateUserRoutes(),
		apihandlers.CreateUserHandler{},
	)

	registerGET[viewhandlers.CreateUserInput](
		context.Background(),
		r,
		viewhandlers.GetCreateUserRoutes(),
		viewhandlers.CreateUserHandler{},
	)
}
