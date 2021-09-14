// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package cmd

import (
	"juno/api/grpc"
	sv "juno/api/grpc/service"
	"juno/api/http"
	"juno/internal/database"
	"juno/internal/endpoint"
	"juno/internal/service"
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
		database.NewMysqlDB,
		service.ProviderSet,
		endpoint.ProviderSet,
		sv.ProviderSet,
		grpc.NewGrpc,
		http.NewHttp,
		NewServer,
	))
}
