package graphql

import (
	"FollowService/internal/domain/entity"

	"github.com/graph-gophers/graphql-go"
)

type UserResolver struct {
	user *entity.User
}

type UserResponseResolver struct {
	user *entity.User
}

func NewUserResolver(user *entity.User) *UserResolver {
	return &UserResolver{user: user}
}

func (r *UserResolver) ID() graphql.ID {
	return graphql.ID(r.user.ID)
}

func (r *UserResolver) Name() string {
	return r.user.Name
}

func (r *UserResponseResolver) Id() graphql.ID {
	return graphql.ID(r.user.ID)
}

func (r *UserResponseResolver) Name() string {
	return r.user.Name
}
