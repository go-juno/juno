package service

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/go-juno/juno/example/juno/api/grpc/protos"
	"github.com/go-juno/juno/example/juno/internal/constant"
	"google.golang.org/grpc"
)

func TestGetGreetingList(t *testing.T) {
	address := fmt.Sprintf("localhost:%d", constant.Config.Server.Grpc.Port)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Panicf("did not connect: %v", err)
	}
	defer conn.Close()
	c := protos.NewGreetingClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	in := &protos.GetGreetingListParam{}
	r, err := c.GetList(ctx, in)
	if err != nil {
		log.Panicf("could not GetGreetingList: %v", err)
	}
	log.Printf("test pass and result:%v", r)
}
