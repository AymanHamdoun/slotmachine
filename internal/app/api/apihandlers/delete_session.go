package apihandlers

import (
	"context"
	"go-backend/internal/app/database/dbrepos"
	"go-backend/internal/app/utils/requestutils"
	"go-backend/internal/app/utils/responseutils"
	"net/http"
)

func GetDeleteSessionRoutes() []string {
	return []string{
		"/api/v1/session",
	}
}

type DeleteSessionInput struct{}

type DeleteSessionResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Deleted bool   `json:"deleted"`
}

type DeleteSessionHandler struct{}

func (h DeleteSessionHandler) Serve(ctx context.Context, input DeleteSessionInput, w http.ResponseWriter, r *http.Request) {
	token, err := requestutils.GetCookieValue(r, "token")
	if err != nil {
		responseutils.WriteJSON(ctx, w, responseutils.ErrorResponse{
			Status:  "unauthenticated",
			Message: "No game session token found",
		})
		return
	}

	repo := dbrepos.NewGameSessionRepo()
	deleted, err := repo.DeleteByToken(ctx, token)
	if err != nil {
		responseutils.WriteJSON(ctx, w, responseutils.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete session",
		})
		return
	}

	if !deleted {
		responseutils.WriteJSON(ctx, w, responseutils.ErrorResponse{
			Status:  "not-found",
			Message: "Session not found or already deleted",
		})
		return
	}

	// Clear the cookie after successful deletion
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    "",
		Path:     "/",
		MaxAge:   -1, // This will delete the cookie
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	})

	responseutils.WriteJSON(ctx, w, DeleteSessionResponse{
		Status:  "ok",
		Message: "Session deleted successfully",
		Deleted: true,
	})
}