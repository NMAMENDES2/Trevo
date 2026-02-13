package api

import (
	"net/http"

	"github.com/NMAMENDES2/Trevo/api/handlers"
	"github.com/NMAMENDES2/Trevo/db"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	router *chi.Mux
	db     *db.Database
}

func NewServer(database *db.Database) *Server {
	s := &Server{
		router: chi.NewRouter(),
		db:     database,
	}

	s.setupMiddleware()
	s.setupRoutes()

	return s
}

func (s *Server) setupMiddleware() {
	s.router.Use(middleware.Logger)
	s.router.Use(middleware.Recoverer)
	s.router.Use(middleware.RequestID)
}

func (s *Server) setupRoutes() {
	userHandler := handlers.NewUserHandler(s.db)
	s.router.Route("/api/v1", func(r chi.Router) {
		r.Get("/users", userHandler.GetUsers)
	})
}

func (s *Server) Start(addr string) error {
	return http.ListenAndServe(addr, s.router)
}
