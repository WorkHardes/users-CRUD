package service

import (
	"context"
	"fmt"

	"github.com/users-CRUD/internal/domain"
	"github.com/users-CRUD/internal/repository"
)

type UserService struct {
	repo   repository.Users
	domain string
}

func NewUserService(repo repository.Users, domain string) *UserService {
	return &UserService{
		repo:   repo,
		domain: domain,
	}
}

func (s *UserService) Create(ctx context.Context, user domain.User) error {
	if err := s.repo.Create(ctx, user); err != nil {
		return fmt.Errorf("s.repo.Create(ctx, user) failed\n%w;", err)
	}
	return nil
}

func (s *UserService) GetById(ctx context.Context, id string) (domain.User, error) {
	user, err := s.repo.GetById(ctx, id)
	if err != nil {
		return domain.User{}, fmt.Errorf("s.repo.GetById(ctx, id) failed\n%w;", err)
	}
	return user, nil
}

func (s *UserService) Update(ctx context.Context, user domain.User) error {
	err := s.repo.Update(ctx, user)
	if err != nil {
		return fmt.Errorf("s.repo.Create(ctx, user) failed\n%w;", err)
	}
	return nil
}

func (s *UserService) Delete(ctx context.Context, id string) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("s.repo.Create(ctx, user) failed\n%w;", err)
	}
	return nil
}
