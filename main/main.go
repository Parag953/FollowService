package main

import (
	"FollowService/Resolver"
	"FollowService/dao"
	"FollowService/handler"
	"github.com/go-chi/chi"
	gql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	dao.AddRandomUsers()
	schemaData, err := ioutil.ReadFile("../config/schema.graphql")
	if err != nil {
		log.Fatalf("Failed to read schema file: %v", err)
	}

	// Parse the schema
	schema := gql.MustParseSchema(string(schemaData), &Resolver.RootResolver{})

	s := handler.Server{Router: chi.NewRouter(), GqlHandler: &handler.GraphiQL{Schema: schema}}
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
