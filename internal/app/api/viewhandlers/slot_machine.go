package viewhandlers

import (
	"context"
	"go-backend/views"
	"net/http"
)

func GetSlotMachineRoutes() []string {
	return []string{
		"/",
		"/slot-machine",
	}
}

type SlotMachineInput struct{}

type SlotMachineHandler struct{}

func (h SlotMachineHandler) Serve(ctx context.Context, input SlotMachineInput, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	views.SlotMachineePage().Render(ctx, w)
}