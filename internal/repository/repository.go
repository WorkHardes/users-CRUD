package repository

import (
	"context"

	"github.com/users-CRUD/internal/domain"
)

type Users interface {
	GetUser(context.Context, domain.User) (*domain.User, error)
	CreateUser(context.Context, domain.User) (*domain.User, error)
	UpdateUser(context.Context, domain.User) (*domain.User, error)
	DeleteUser(context.Context, domain.User) error
}

type Repositories struct {
	Users Users
}

func NewRepositories() string {
	return ""
}
