package graphql

import (
	"FollowService/internal/domain/service"
	"context"

	"github.com/graph-gophers/graphql-go"
)

type RootResolver struct {
	userService *service.UserService
}

func NewRootResolver(userService *service.UserService) *RootResolver {
	return &RootResolver{
		userService: userService,
	}
}

func (r *RootResolver) CreateUser(ctx context.Context, args struct{ Name string }) (*UserResponseResolver, error) {
	user, err := r.userService.CreateUser(ctx, args.Name)
	if err != nil {
		return nil, err
	}
	return &UserResponseResolver{user: user}, nil
}

func (r *RootResolver) FollowUser(ctx context.Context, args struct{ MyId, TargetId graphql.ID }) (bool, error) {
	err := r.userService.FollowUser(ctx, string(args.MyId), string(args.TargetId))
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *RootResolver) UnfollowUser(ctx context.Context, args struct{ MyId, TargetId graphql.ID }) (bool, error) {
	err := r.userService.UnfollowUser(ctx, string(args.MyId), string(args.TargetId))
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *RootResolver) Followers(ctx context.Context, args struct{ Id graphql.ID }) ([]*UserResponseResolver, error) {
	users, err := r.userService.GetFollowers(ctx, string(args.Id))
	if err != nil {
		return nil, err
	}

	var resolvers []*UserResponseResolver
	for _, user := range users {
		resolvers = append(resolvers, &UserResponseResolver{user: user})
	}
	return resolvers, nil
}

func (r *RootResolver) Followings(ctx context.Context, args struct{ Id graphql.ID }) ([]*UserResponseResolver, error) {
	users, err := r.userService.GetFollowing(ctx, string(args.Id))
	if err != nil {
		return nil, err
	}

	var resolvers []*UserResponseResolver
	for _, user := range users {
		resolvers = append(resolvers, &UserResponseResolver{user: user})
	}
	return resolvers, nil
}
