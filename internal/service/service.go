package service

import (
	"context"

	"github.com/users-CRUD/internal/domain"
	"github.com/users-CRUD/internal/repository"
)

type Users interface {
	Create(ctx context.Context, user domain.User) error
	GetById(ctx context.Context, id string) (domain.User, error)
	Update(ctx context.Context, user domain.User) error
	Delete(ctx context.Context, id string) error
}
type Services struct {
	Users Users
}
type Deps struct {
	Repos  *repository.Repositories
	Domain string
}

func NewServices(depts Deps) *Services {
	usersService := NewUserService(depts.Repos.Users, depts.Domain)

	return &Services{
		Users: usersService,
	}
}
