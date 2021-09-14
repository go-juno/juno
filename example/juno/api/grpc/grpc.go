package grpc

import (
	"github.com/go-juno/juno/example/juno/api/grpc/protos"
	"github.com/go-juno/juno/example/juno/api/grpc/service"
	"google.golang.org/grpc"
)

func NewGrpc(
	greeting *service.GreetingServer,
) (s *grpc.Server, err error) {
	s = grpc.NewServer()
	protos.RegisterGreetingServer(s, greeting)
	return
}
