package viewhandlers

import (
	"context"
	"go-backend/views"
	"net/http"
)

func GetMinimalSlotRoutes() []string {
	return []string{
		"/minimal",
		"/minimal-slot",
	}
}

type MinimalSlotInput struct{}

type MinimalSlotHandler struct{}

func (h MinimalSlotHandler) Serve(ctx context.Context, input MinimalSlotInput, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	views.MinimalSlotPage().Render(ctx, w)
}