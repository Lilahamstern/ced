package resolver

import (
	"github.com/lilahamstern/ced/server/pkg/service"
)

//go:generate go run github.com/99designs/gqlgen

var resolver *Resolver

type Resolver struct {
	services *service.Services
}

func NewResolver(services *service.Services) {
	resolver = &Resolver{
		services,
	}
}
