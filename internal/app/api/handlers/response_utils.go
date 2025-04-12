package handlers

import (
	"context"
	"encoding/json"
	"net/http"
)

func _json(ctx context.Context, w http.ResponseWriter, structToWrite any) {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(structToWrite)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
