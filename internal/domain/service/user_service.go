package service

import (
	"FollowService/internal/domain/entity"
	"FollowService/internal/domain/repository"
	"context"
	"errors"
	"fmt"
	"time"
)

type UserService struct {
	userRepo   repository.UserRepository
	followRepo repository.FollowRepository
}

func NewUserService(userRepo repository.UserRepository, followRepo repository.FollowRepository) *UserService {
	return &UserService{
		userRepo:   userRepo,
		followRepo: followRepo,
	}
}

func (s *UserService) CreateUser(ctx context.Context, name string) (*entity.User, error) {
	user := &entity.User{
		ID:   fmt.Sprintf("user_%d", time.Now().Unix()),
		Name: name,
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) FollowUser(ctx context.Context, followerID, followeeID string) error {
	if followerID == followeeID {
		return errors.New("cannot follow yourself")
	}

	// Check if users exist
	if _, err := s.userRepo.GetByID(ctx, followerID); err != nil {
		return err
	}
	if _, err := s.userRepo.GetByID(ctx, followeeID); err != nil {
		return err
	}

	// Check if already following
	isFollowing, err := s.followRepo.IsFollowing(ctx, followerID, followeeID)
	if err != nil {
		return err
	}
	if isFollowing {
		return errors.New("already following this user")
	}

	return s.followRepo.Follow(ctx, followerID, followeeID)
}

func (s *UserService) UnfollowUser(ctx context.Context, followerID, followeeID string) error {
	if followerID == followeeID {
		return errors.New("cannot unfollow yourself")
	}

	// Check if users exist
	if _, err := s.userRepo.GetByID(ctx, followerID); err != nil {
		return err
	}
	if _, err := s.userRepo.GetByID(ctx, followeeID); err != nil {
		return err
	}

	// Check if currently following
	isFollowing, err := s.followRepo.IsFollowing(ctx, followerID, followeeID)
	if err != nil {
		return err
	}
	if !isFollowing {
		return errors.New("not following this user")
	}

	return s.followRepo.Unfollow(ctx, followerID, followeeID)
}

func (s *UserService) GetFollowers(ctx context.Context, userID string) ([]*entity.User, error) {
	// Check if user exists
	if _, err := s.userRepo.GetByID(ctx, userID); err != nil {
		return nil, err
	}

	return s.followRepo.GetFollowers(ctx, userID)
}

func (s *UserService) GetFollowing(ctx context.Context, userID string) ([]*entity.User, error) {
	// Check if user exists
	if _, err := s.userRepo.GetByID(ctx, userID); err != nil {
		return nil, err
	}

	return s.followRepo.GetFollowing(ctx, userID)
}
