package handler

import (
	"github.com/go-chi/chi"
)

type Server struct {
	Router     chi.Router
	GqlHandler *GraphiQL
}
