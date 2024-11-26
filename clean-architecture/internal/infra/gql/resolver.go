package gql

//go:generate go run github.com/99designs/gqlgen generate

import (
	"github.com/HenriqueFigueiredo1/clean-architecture/internal/usecase"
)

type Resolver struct {
	CreateOrderUseCase usecase.CreateOrderUseCase
	ListOrderUseCase   usecase.ListOrderUseCase
}
