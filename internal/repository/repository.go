package repository

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/users-CRUD/internal/domain"
)

type Users interface {
	Create(ctx context.Context, user domain.User) error
	GetById(ctx context.Context, id string) (domain.User, error)
	Update(ctx context.Context, user domain.User) error
	Delete(ctx context.Context, id string) error
}

type Repositories struct {
	Users Users
}

func NewRepositories(db *pgx.Conn) *Repositories {
	return &Repositories{
		Users: NewUserRepo(db),
	}
}
