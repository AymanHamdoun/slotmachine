package viewhandlers

import (
	"context"
	"net/http"
)

func GetCreateUserRoutes() []string {
	return []string{
		"/users/create",
	}
}

type CreateUserInput struct{}

type CreateUserHandler struct{}

type createUserViewData struct {
}

func (h CreateUserHandler) Serve(ctx context.Context, input CreateUserInput, w http.ResponseWriter, r *http.Request) {
	_view(ctx, w, r, "index.html", createUserViewData{})
}
