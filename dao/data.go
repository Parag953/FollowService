package dao

import (
	"FollowService/model"
	"fmt"
	gql "github.com/graph-gophers/graphql-go"
)

var Users []*model.User

func createRandomUser(id int) *model.User {
	return &model.User{
		Id:         gql.ID(fmt.Sprintf("user%d", id)),
		Name:       fmt.Sprintf("User %d", id),
		Followers:  make([]gql.ID, 0),
		Followings: make([]gql.ID, 0),
	}
}

func AddRandomUsers() {

	for i := 1; i <= 5; i++ {
		Users = append(Users, createRandomUser(i))
	}

	Users[0].Followers = append(Users[0].Followers, Users[1].Id, Users[2].Id)
	Users[1].Followings = append(Users[1].Followings, Users[0].Id)
	Users[2].Followings = append(Users[2].Followings, Users[0].Id)

	// Print the users to verify
	for _, user := range Users {
		fmt.Printf("ID: %s, Name: %s\n", user.Id, user.Name)
	}
}
