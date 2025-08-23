package repository

import (
	"FollowService/internal/domain/entity"
	"context"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) error
	GetByID(ctx context.Context, id string) (*entity.User, error)
	GetAll(ctx context.Context) ([]*entity.User, error)
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, id string) error
}

type FollowRepository interface {
	Follow(ctx context.Context, followerID, followeeID string) error
	Unfollow(ctx context.Context, followerID, followeeID string) error
	GetFollowers(ctx context.Context, userID string) ([]*entity.User, error)
	GetFollowing(ctx context.Context, userID string) ([]*entity.User, error)
	IsFollowing(ctx context.Context, followerID, followeeID string) (bool, error)
}
