package apihandlers

import (
	"context"
	"go-backend/internal/app/database/dbrepos"
	"go-backend/internal/app/utils/requestutils"
	"go-backend/internal/app/utils/responseutils"
	"net/http"
)

func GetGetSessionRoutes() []string {
	return []string{
		"/api/v1/session",
	}
}

type GetSessionInput struct {
	Name string `json:"name" form:"name"`
}

type GetSessionResponse struct {
	Status  string              `json:"status"`
	Session GameSessionResource `json:"session"`
}

type GetSessionHandler struct{}

func (h GetSessionHandler) Serve(ctx context.Context, input GetSessionInput, w http.ResponseWriter, r *http.Request) {
	token, err := requestutils.GetCookieValue(r, "token")
	if err != nil {
		responseutils.WriteJSON(ctx, w, responseutils.ErrorResponse{
			Status:  "unauthenticated",
			Message: "No game session token found",
		})
		return
	}

	repo := dbrepos.NewGameSessionRepo()
	var session dbrepos.GameSession
	session, err = repo.GetByToken(ctx, token)
	if err != nil {
		responseutils.WriteJSON(ctx, w, responseutils.ErrorResponse{
			Status:  "invalid-session",
			Message: "Game session is not valid",
		})
		return
	}

	responseutils.WriteJSON(ctx, w, GetSessionResponse{
		Status: "ok",
		Session: GameSessionResource{
			Name:    session.Name,
			Token:   session.Token,
			Credits: session.Credits,
		},
	})
}
