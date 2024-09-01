package model

import "github.com/graph-gophers/graphql-go"

type User struct {
	Id         graphql.ID   `json:"id"`
	Name       string       `json:"name"`
	Followers  []graphql.ID `json:"followers"`
	Followings []graphql.ID `json:"followings"`
}

type UserResponse struct {
	Id   graphql.ID `json:"id"`
	Name string     `json:"name"`
}
