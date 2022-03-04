package server

import (
	"context"
	"net/http"
	"time"

	"github.com/users-CRUD/internal/config"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg *config.Config, handler http.Handler) *Server {
	s := &http.Server{
		Addr:           ":" + cfg.HTTP.Port,
		Handler:        handler,
		ReadTimeout:    cfg.HTTP.ReadTimeout * time.Second,
		WriteTimeout:   cfg.HTTP.WriteTimeout * time.Second,
		MaxHeaderBytes: cfg.HTTP.MaxHeaderBytes,
	}
	return &Server{
		httpServer: s,
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
