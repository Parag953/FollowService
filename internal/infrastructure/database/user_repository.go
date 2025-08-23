package database

import (
	"FollowService/internal/domain/entity"
	"FollowService/internal/domain/repository"
	"context"
	"database/sql"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *entity.User) error {
	query := `INSERT INTO users (id, name, created_at, updated_at) VALUES ($1, $2, NOW(), NOW())`
	_, err := r.db.ExecContext(ctx, query, user.ID, user.Name)
	return err
}

func (r *userRepository) GetByID(ctx context.Context, id string) (*entity.User, error) {
	query := `SELECT id, name, created_at, updated_at FROM users WHERE id = $1`
	user := &entity.User{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Name, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) GetAll(ctx context.Context) ([]*entity.User, error) {
	query := `SELECT id, name, created_at, updated_at FROM users`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []*entity.User
	for rows.Next() {
		user := &entity.User{}
		if err := rows.Scan(&user.ID, &user.Name, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *userRepository) Update(ctx context.Context, user *entity.User) error {
	query := `UPDATE users SET name = $1, updated_at = NOW() WHERE id = $2`
	_, err := r.db.ExecContext(ctx, query, user.Name, user.ID)
	return err
}

func (r *userRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
