package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	router *chi.Mux
}

func NewServer() *Server {
	r := chi.NewRouter()

	// Basic middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)

	return &Server{
		router: r,
	}
}

func (s *Server) Router() *chi.Mux {
	return s.router
}

func (s *Server) Start(addr string) error {
	return http.ListenAndServe(addr, s.router)
}
