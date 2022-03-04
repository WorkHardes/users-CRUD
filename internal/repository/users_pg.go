package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/users-CRUD/internal/domain"
)

type UsersRepo struct {
	db *pgx.Conn
}

func NewUserRepo(db *pgx.Conn) *UsersRepo {
	return &UsersRepo{
		db: db,
	}
}

func (r *UsersRepo) GetById(ctx context.Context, id string) (domain.User, error) {
	var user domain.User
	query := fmt.Sprintf("SELECT name, email, password FROM users WHERE id=%s", id)
	err := r.db.QueryRow(ctx, query).Scan(&user.Name, &user.Email, &user.Password)
	if err != nil {
		return domain.User{}, fmt.Errorf("UsersRepo.GetById(): conn.QueryRow failed\n%w;", err)
	}
	return user, nil
}

func (r *UsersRepo) Create(ctx context.Context, user domain.User) error {
	query := fmt.Sprintf("INSERT INTO users(name, email, password) VALUES('%s', '%s', '%s') returning id;", user.Name, user.Email, user.Password)
	var id int64
	err := r.db.QueryRow(ctx, query).Scan(&id)
	if err != nil {
		return fmt.Errorf("UsersRepo.Create(): conn.QueryRow failed;%w;", err)
	}
	return nil
}

func (r *UsersRepo) Update(ctx context.Context, user domain.User) error {
	query := fmt.Sprintf("UPDATE users SET name='%s', email='%s', password='%s' WHERE id='%s';", user.Name, user.Email, user.Password, user.ID)
	_, err := r.db.Query(ctx, query)
	if err != nil {
		return fmt.Errorf("UsersRepo.Update(): conn.QueryRow failed;%w;", err)
	}
	return nil
}

func (r *UsersRepo) Delete(ctx context.Context, id string) error {
	query := fmt.Sprintf("DELETE FROM users WHERE id='%s';", id)
	_, err := r.db.Query(ctx, query)
	if err != nil {
		return fmt.Errorf("UsersRepo.Delete(): conn.QueryRow failed;%w;", err)
	}
	return nil
}
