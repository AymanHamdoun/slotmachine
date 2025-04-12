package handlers

import (
	"context"
	"net/http"
)

func GetCreateUserRoutes() []string {
	return []string{
		"/api/v1/users",
	}
}

type CreateUserInput struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type CreateUserResponse struct {
	Status string `json:"status"`
}

type CreateUserHandler struct{}

func (h CreateUserHandler) Serve(ctx context.Context, input CreateUserInput, w http.ResponseWriter) {
	_json(ctx, w, CreateUserResponse{
		Status: "ok",
	})
}
