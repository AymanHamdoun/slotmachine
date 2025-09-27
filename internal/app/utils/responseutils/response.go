package responseutils

import (
	"context"
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func WriteJSON(ctx context.Context, w http.ResponseWriter, structToWrite any) {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(structToWrite)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func SetCookie(w http.ResponseWriter, httpOnly bool, key string, value string) {
	http.SetCookie(w, &http.Cookie{
		Name:     key,
		Value:    value,
		Path:     "/",
		HttpOnly: httpOnly,             // not accessible to JS.
		Secure:   true,                 // only over HTTPS.
		SameSite: http.SameSiteLaxMode, // prevents browser from sending cooking to other sites.
	})
}
