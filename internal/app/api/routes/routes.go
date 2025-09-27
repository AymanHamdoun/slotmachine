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

	// Serves static files like custom JS, CSS, etc.
	staticFS := http.FileServer(http.Dir("web/static"))
	r.Handle("/static/*", http.StripPrefix("/static/", staticFS))

	registerPOST[apihandlers.CreateSessionInput](
		context.Background(),
		r,
		apihandlers.GetCreateSessionRoutes(),
		apihandlers.CreateSessionHandler{},
	)

	registerPOST[apihandlers.RollInput](
		context.Background(),
		r,
		apihandlers.GetRollRoutes(),
		apihandlers.RollHandler{},
	)

	registerPOST[apihandlers.CashOutInput](
		context.Background(),
		r,
		apihandlers.GetCashOutRoutes(),
		apihandlers.CashOutHandler{},
	)

	registerGET[apihandlers.GetSessionInput](
		context.Background(),
		r,
		apihandlers.GetGetSessionRoutes(),
		apihandlers.GetSessionHandler{},
	)

	registerGET[viewhandlers.SlotMachineInput](
		context.Background(),
		r,
		viewhandlers.GetSlotMachineRoutes(),
		viewhandlers.SlotMachineHandler{},
	)

	registerGET[viewhandlers.MinimalSlotInput](
		context.Background(),
		r,
		viewhandlers.GetMinimalSlotRoutes(),
		viewhandlers.MinimalSlotHandler{},
	)
}
