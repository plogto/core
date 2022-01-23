package graph

import "github.com/plogto/core/service"

//go:generate go run github.com/99designs/gqlgen
type Resolver struct {
	Service *service.Service
}
