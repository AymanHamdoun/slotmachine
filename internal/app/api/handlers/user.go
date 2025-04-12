package handlers

import (
	"context"
	"database/sql"
	"go-backend/internal/app/database"
	"go-backend/internal/app/database/dbmodels"
	"net/http"
)

func GetCreateUserRoutes() []string {
	return []string{
		"/api/v1/users",
	}
}

type CreateUserInput struct {
	FirstName string `json:"first_name" form:"first_name"`
	LastName  string `json:"last_name" form:"last_name"`
	Email     string `json:"email" form:"email"`
	Password  string `json:"password" form:"password"`
}

type CreateUserResponse struct {
	Status string `json:"status"`
}

type CreateUserHandler struct{}

func (h CreateUserHandler) Serve(ctx context.Context, input CreateUserInput, w http.ResponseWriter) {
	db, err := database.GetDB(ctx)
	if err != nil {
		_json(ctx, w, errorResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	queries := dbmodels.New(db)

	queries.CreateUserWithEmail(ctx, dbmodels.CreateUserWithEmailParams{
		FirstName:            input.FirstName,
		LastName:             input.LastName,
		Email:                sql.NullString{String: input.Email},
		RegistrationMethodID: 1,
	})

	_json(ctx, w, CreateUserResponse{
		Status: "ok",
	})
}
