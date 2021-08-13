package graph

import "github.com/favecode/note-core/service"

//go:generate go run github.com/99designs/gqlgen
type Resolver struct {
	Service *service.Service
}
