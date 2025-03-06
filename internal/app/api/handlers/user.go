package handlers

import (
	"encoding/json"
	"net/http"

	"go-backend/internal/app/api/middleware"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

type CreateUserInput struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input CreateUserInput

	if err := middleware.BindInputData(r, &input); err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	// Here you would typically:
	// 1. Validate the input
	// 2. Hash the password
	// 3. Save to database
	// 4. Return response

	// For now, we'll just return the received data
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User created successfully",
		"user": map[string]string{
			"name":  input.Name,
			"email": input.Email,
		},
	})
}
