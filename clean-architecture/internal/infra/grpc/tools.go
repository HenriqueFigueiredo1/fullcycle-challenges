package grpc

//go:generate protoc --go_out=. --go-grpc_out=. ./protofiles/order.proto

import (
	_ "github.com/HenriqueFigueiredo1/clean-architecture/internal/infra/grpc/pb"
	_ "github.com/HenriqueFigueiredo1/clean-architecture/internal/infra/grpc/service"
)
