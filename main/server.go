package main

import (
	"FollowService/handler"
	"github.com/go-chi/chi"
)

type Server struct {
	Router     chi.Router
	GqlHandler *handler.GraphiQL
}
