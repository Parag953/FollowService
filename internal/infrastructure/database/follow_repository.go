package database

import (
	"FollowService/internal/domain/entity"
	"FollowService/internal/domain/repository"
	"context"
	"database/sql"
)

type followRepository struct {
	db *sql.DB
}

func NewFollowRepository(db *sql.DB) repository.FollowRepository {
	return &followRepository{db: db}
}

func (r *followRepository) Follow(ctx context.Context, followerID, followeeID string) error {
	query := `INSERT INTO follows (follower_id, followee_id) VALUES (?, ?)`
	_, err := r.db.ExecContext(ctx, query, followerID, followeeID)
	return err
}

func (r *followRepository) Unfollow(ctx context.Context, followerID, followeeID string) error {
	query := `DELETE FROM follows WHERE follower_id = ? AND followee_id = ?`
	_, err := r.db.ExecContext(ctx, query, followerID, followeeID)
	return err
}

func (r *followRepository) GetFollowers(ctx context.Context, userID string) ([]*entity.User, error) {
	query := `
		SELECT u.id, u.name, u.created_at, u.updated_at 
		FROM users u 
		INNER JOIN follows f ON u.id = f.follower_id 
		WHERE f.followee_id = ?`

	rows, err := r.db.QueryContext(ctx, query, userID)
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

func (r *followRepository) GetFollowing(ctx context.Context, userID string) ([]*entity.User, error) {
	query := `
		SELECT u.id, u.name, u.created_at, u.updated_at 
		FROM users u 
		INNER JOIN follows f ON u.id = f.followee_id 
		WHERE f.follower_id = ?`

	rows, err := r.db.QueryContext(ctx, query, userID)
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

func (r *followRepository) IsFollowing(ctx context.Context, followerID, followeeID string) (bool, error) {
	query := `SELECT COUNT(*) FROM follows WHERE follower_id = ? AND followee_id = ?`
	var count int
	err := r.db.QueryRowContext(ctx, query, followerID, followeeID).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
