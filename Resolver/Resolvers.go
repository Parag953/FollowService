package main

import (
	"errors"
	"fmt"
	"github.com/graph-gophers/graphql-go"
)

type args struct {
	MyId     graphql.ID
	TargetId graphql.ID
}

type RootResolver struct{}

type UserResolver struct{ u *User }

type UserResponseResolver struct{ u *UserResponse }

func (r *RootResolver) Users() ([]*UserResolver, error) {
	var userRxs []*UserResolver
	for _, u := range Users {
		userRxs = append(userRxs, &UserResolver{u})
	}
	return userRxs, nil
}

func (r *RootResolver) User(args struct{ Id graphql.ID }) (*UserResolver, error) {
	for _, u := range Users {
		if u.Id == args.Id {
			return &UserResolver{u}, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("user with Id: %v not found", args.Id))
}

func (r *RootResolver) Followers(args struct{ Id graphql.ID }) ([]*UserResponseResolver, error) {
	user, err := r.User(args)
	if err != nil {
		return nil, err
	}
	var followerRxs []*UserResponseResolver
	for _, followerId := range user.u.Followers {

		arg := struct{ Id graphql.ID }{followerId}
		userResolver, _ := r.User(arg)
		name := userResolver.u.Name
		followerRxs = append(followerRxs, &UserResponseResolver{&UserResponse{followerId, name}})
	}
	return followerRxs, nil
}

func (r *RootResolver) Followings(args struct{ Id graphql.ID }) ([]*UserResponseResolver, error) {
	user, err := r.User(args)
	if err != nil {
		return nil, err
	}
	var followingRxs []*UserResponseResolver
	for _, followingId := range user.u.Followings {

		arg := struct{ Id graphql.ID }{followingId}
		userResolver, _ := r.User(arg)
		name := userResolver.u.Name
		followingRxs = append(followingRxs, &UserResponseResolver{&UserResponse{followingId, name}})
	}
	return followingRxs, nil
}

func (r *UserResolver) Id() graphql.ID {
	return r.u.Id
}
func (r *UserResolver) Name() string {
	return r.u.Name
}

func (r *UserResolver) Followers() []graphql.ID {
	return r.u.Followers
}

func (r *UserResolver) Followeings() []graphql.ID {
	return r.u.Followings
}

func (r *UserResponseResolver) Id() graphql.ID {
	return r.u.Id
}

func (r *UserResponseResolver) Name() string {
	return r.u.Name
}

func (r *RootResolver) FollowUser(args args) (bool, error) {

	if args.MyId == args.TargetId {
		return false, errors.New("you can't follow yourself")
	}

	myUser, err := r.User(struct{ Id graphql.ID }{args.MyId})
	if err != nil {
		return false, err
	}

	targetUser, err := r.User(struct{ Id graphql.ID }{args.TargetId})
	if err != nil {
		return false, err
	}

	// for loop for checking if the user is already following the target
	for _, followingId := range myUser.u.Followings {
		if followingId == args.TargetId {
			return false, errors.New("you are already following this user")
		}
	}

	myUser.u.Followings = append(myUser.u.Followings, args.TargetId)
	targetUser.u.Followers = append(targetUser.u.Followers, args.MyId)

	return true, nil
}

func (r *RootResolver) UnfollowUser(args args) (bool, error) {

	if args.MyId == args.TargetId {
		return false, errors.New("you can't unfollow yourself")
	}

	myUser, err := r.User(struct{ Id graphql.ID }{args.MyId})
	if err != nil {
		return false, err
	}
	for i, followingId := range myUser.u.Followings {
		if followingId == args.TargetId {
			myUser.u.Followings = append(myUser.u.Followings[:i], myUser.u.Followings[i+1:]...)
			break
		}
		if i == len(myUser.u.Followings)-1 {
			return false, errors.New("you are not following this user")
		}
	}

	targetUser, err := r.User(struct{ Id graphql.ID }{args.TargetId})
	if err != nil {
		return false, err
	}
	for i, followerId := range targetUser.u.Followers {
		if followerId == args.MyId {
			targetUser.u.Followers = append(targetUser.u.Followers[:i], targetUser.u.Followers[i+1:]...)
			break
		}
	}
	return true, nil
}

func (r *RootResolver) CreateUser(args struct{ Name string }) (*UserResponseResolver, error) {
	newUser := &User{
		Id:         graphql.ID(fmt.Sprintf("user%d", len(Users)+1)),
		Name:       args.Name,
		Followers:  make([]graphql.ID, 0),
		Followings: make([]graphql.ID, 0),
	}
	Users = append(Users, newUser)
	u := &UserResponse{
		Id:   newUser.Id,
		Name: newUser.Name,
	}
	return &UserResponseResolver{u}, nil
}
