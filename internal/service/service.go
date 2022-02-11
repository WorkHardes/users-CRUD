package service

import "github.com/users-CRUD/internal/repository"

type Deps struct {
	Repo   repository.Repositories
	Domain string
}

func NewServices() string {
	return ""
}
