package http

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/users-CRUD/internal/config"
	delivery "github.com/users-CRUD/internal/delivery/http"
	"github.com/users-CRUD/internal/repository"
	"github.com/users-CRUD/internal/server"
	"github.com/users-CRUD/internal/service"
	"github.com/users-CRUD/pkg/database/pgdb"
	"github.com/users-CRUD/pkg/logger"
)

func Run(configsDir string) {
	cfg, err := config.Init(configsDir)
	if err != nil {
		logger.Error(err)
		return
	}

	postgresClient, err := pgdb.NewClient(cfg.Postgres.URI, cfg.Postgres.User, cfg.Postgres.Password)
	if err != nil {
		logger.Error(err)
		return
	}

	db := postgresClient.Database(cfg.Postgres.Name)

	repo := repository.NewRepositories(db)
	services := service.NewServices(service.Deps{
		Repo:   repo,
		Domain: cfg.HTTP.Host,
	})

	handlers := delivery.NewHandler(services)
	srv := server.NewServer(cfg, handlers)

	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			logger.Errorf("error occurred while running http server: %s\n", err.Error())
		}
	}()

	logger.Info("Server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 5 * time.Second
	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		logger.Errorf("failed to stop server: %v", err)
	}
}
