package repository

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/users-CRUD/internal/config"
)

func RunPgMigrations(cfg *config.PostgresConfig) error {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.UserName, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
	if cfg.MigrationsPath == "" {
		return fmt.Errorf("config.PostgresConfig.MigrationsPath == ''")
	}

	m, err := migrate.New(
		cfg.MigrationsPath,
		dbURL,
	)
	if err != nil {
		return fmt.Errorf("migrate.New failed\n%w;", err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("migrate.New.Up failed\n%w;", err)
	}
	return nil
}
