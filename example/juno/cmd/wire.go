//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package cmd

import (
	"github.com/go-juno/juno/example/juno/api/grpc"
	sv "github.com/go-juno/juno/example/juno/api/grpc/service"
	"github.com/go-juno/juno/example/juno/api/http"
	"github.com/go-juno/juno/example/juno/internal/endpoint"
	"github.com/go-juno/juno/example/juno/internal/repo"
	"github.com/go-juno/juno/example/juno/internal/service"
	"github.com/google/wire"
	http2 "net/http"

	grpc2 "google.golang.org/grpc"
)

type Server struct {
	Http *http2.Server
	Grpc *grpc2.Server
}

func NewServer(http *http2.Server,
	grpc *grpc2.Server) *Server {
	return &Server{
		Grpc: grpc,
		Http: http,
	}
}

// initApp init endpoints
func InitServer() (*Server, error) {
	panic(wire.Build(
		repo.ProviderSet,
		service.ProviderSet,
		endpoint.ProviderSet,
		sv.ProviderSet,
		grpc.NewGrpc,
		http.NewHttp,
		NewServer,
	))
}
