package apihandlers

import (
	"context"
	"go-backend/internal/app/database/dbrepos"
	"go-backend/internal/app/utils/requestutils"
	"go-backend/internal/app/utils/responseutils"
	"net/http"
)

func GetCashOutRoutes() []string {
	return []string{
		"/api/v1/cash-out",
	}
}

type CashOutInput struct {
	AccountNumber string `json:"account_number" form:"account_number"`
}

type CashOutResponse struct {
	Status        string `json:"status"`
	Message       string `json:"message"`
	CashedOut     int32  `json:"cashed_out"`
	AccountNumber string `json:"account_number"`
}

type CashOutHandler struct{}

func (h CashOutHandler) Serve(ctx context.Context, input CashOutInput, w http.ResponseWriter, r *http.Request) {
	// Get token from cookie
	token, err := requestutils.GetCookieValue(r, "token")
	if err != nil {
		responseutils.WriteJSON(ctx, w, responseutils.ErrorResponse{
			Status:  "unauthenticated",
			Message: "No game session token found",
		})
		return
	}

	// Validate account number is provided
	if input.AccountNumber == "" {
		responseutils.WriteJSON(ctx, w, responseutils.ErrorResponse{
			Status:  "invalid-input",
			Message: "Account number is required",
		})
		return
	}

	repo := dbrepos.NewGameSessionRepo()

	// Get the current session to know the credits amount
	session, err := repo.GetByToken(ctx, token)
	if err != nil {
		responseutils.WriteJSON(ctx, w, responseutils.ErrorResponse{
			Status:  "invalid-session",
			Message: "Game session is not valid",
		})
		return
	}

	creditsToReturn := session.Credits

	// Set credits to 0 first (this ensures we don't lose track if delete fails)
	// We subtract all the credits the user has
	if session.Credits > 0 {
		success, err := repo.SubtractCredits(ctx, token, session.Credits)
		if err != nil || !success {
			responseutils.WriteJSON(ctx, w, responseutils.ErrorResponse{
				Status:  "error",
				Message: "Failed to reset credits",
			})
			return
		}
	}

	// Delete the game session (soft delete)
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

	// Clear the session cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    "",
		Path:     "/",
		MaxAge:   -1, // This will delete the cookie
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	})

	responseutils.WriteJSON(ctx, w, CashOutResponse{
		Status:        "ok",
		Message:       "Successfully cashed out",
		CashedOut:     creditsToReturn,
		AccountNumber: input.AccountNumber,
	})
}
