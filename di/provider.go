// +build wireinject

package di

import (
	"github.com/google/wire"
	"grpc-golang/handler"
	"google.golang.org/grpc"
)

var providerSet = wire.NewSet(
	handler.NewHelloQuery,
	handler.NewAPIHandler,
)

func ResolveAPIHandler() *grpc.Server {
	panic(wire.Build(providerSet))
}
