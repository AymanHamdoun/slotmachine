package apihandlers

import (
	"context"
	"go-backend/internal/app/database/dbrepos"
	"net/http"

	"github.com/google/uuid"
)

func GetCreateSessionRoutes() []string {
	return []string{
		"/api/v1/sessions",
	}
}

type CreateSessionInput struct {
	Name string `json:"name" form:"name"`
}

type GameSessionResource struct {
	Name    string `json:"name"`
	Token   string `json:"token"`
	Credits int32  `json:"credits"`
}

type CreateSessionResponse struct {
	Status  string              `json:"status"`
	Session GameSessionResource `json:"session"`
}

type CreateSessionHandler struct{}

func (h CreateSessionHandler) Serve(ctx context.Context, input CreateSessionInput, w http.ResponseWriter, r *http.Request) {
	const defaultCredits = 10
	sessionToken := uuid.New().String()

	gameSessionRepo := dbrepos.NewGameSessionRepo()
	session, err := gameSessionRepo.Create(ctx, dbrepos.GameSessionParams{
		Name:    input.Name,
		Token:   sessionToken,
		Credits: defaultCredits,
	})

	if err != nil {
		_json(ctx, w, errorResponse{
			Status:  "error",
			Message: err.Error(),
		})
	}

	_json(ctx, w, CreateSessionResponse{
		Status: "ok",
		Session: GameSessionResource{
			Name:    session.Name,
			Token:   session.Token,
			Credits: session.Credits,
		},
	})
}
