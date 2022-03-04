package pgdb

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/users-CRUD/internal/config"
)

// const timeout = 10 * time.Second

// pg client imports from library
func NewClient(ctx context.Context, cfg *config.PostgresConfig) (*pgx.Conn, error) {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.UserName, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
	dbConn, err := pgx.Connect(ctx, dbURL)
	if err != nil {
		return nil, fmt.Errorf("pgx.Connect failed; %w; ", err)
	}

	return dbConn, nil
}
