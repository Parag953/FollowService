package main

import (
	"FollowService/handler"
	"fmt"
	"github.com/go-chi/chi"
	gql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"io/ioutil"
	"log"
	"net/http"
)

var Users []*User

func createRandomUser(id int) *User {
	return &User{
		Id:         gql.ID(fmt.Sprintf("user%d", id)),
		Name:       fmt.Sprintf("User %d", id),
		Followers:  make([]gql.ID, 0),
		Followings: make([]gql.ID, 0),
	}
}

func addRandomUsers() {

	for i := 1; i <= 5; i++ {
		Users = append(Users, createRandomUser(i))
	}

	Users[0].Followers = append(Users[0].Followers, Users[1].Id, Users[2].Id)
	Users[0].Followings = append(Users[0].Followings, Users[1].Id, Users[3].Id)

	// Print the users to verify
	for _, user := range Users {
		fmt.Printf("ID: %s, Name: %s\n", user.Id, user.Name)
	}
}

func main() {
	addRandomUsers()
	schemaData, err := ioutil.ReadFile("schema.graphql")
	if err != nil {
		log.Fatalf("Failed to read schema file: %v", err)
	}

	// Parse the schema
	schema := gql.MustParseSchema(string(schemaData), &RootResolver{})

	s := Server{Router: chi.NewRouter(), GqlHandler: &handler.GraphiQL{Schema: schema}}
	// Set up the HTTP handler
	s.Router.Group(func(r chi.Router) {
		r.Handle("/query", &relay.Handler{Schema: s.GqlHandler.Schema})
		r.Get("/graphiql", func(w http.ResponseWriter, r *http.Request) {
			w.Write(handler.Page)
		})
	})

	log.Println("Server is running on http://localhost:8080/graphiql")
	log.Fatal(http.ListenAndServe(":8080", s.Router))
}
