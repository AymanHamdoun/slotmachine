package routes

import (
	"context"
	"go-backend/internal/app/api/apihandlers"
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

	registerPOST[apihandlers.CreateSessionInput](
		context.Background(),
		r,
		apihandlers.GetCreateSessionRoutes(),
		apihandlers.CreateSessionHandler{},
	)

	registerGET[apihandlers.GetSessionInput](
		context.Background(),
		r,
		apihandlers.GetGetSessionRoutes(),
		apihandlers.GetSessionHandler{},
	)
}
