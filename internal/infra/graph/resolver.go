package graph

import "github.com/andreflor21/go-clean/internal/usecase"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	CreateOrderUseCase usecase.CreateOrderUseCase
}
