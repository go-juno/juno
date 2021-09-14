package cmd

import (
	"context"
	"fmt"
	"log"
	"net"

	"juno/internal/constant"

	"golang.org/x/sync/errgroup"
)

func Start() {
	log.SetFlags(log.Llongfile | log.LstdFlags)

	g, _ := errgroup.WithContext(context.Background())
	server, err := InitServer()
	if err != nil {
		panic(err)
	}

	// 启动http 服务
	g.Go(func() error {
		log.Printf(" [*] http serve Waiting for connection:%d", constant.Config.Server.Http.Port)
		err := server.Http.ListenAndServe()
		return err
	})

	//启动grpc服务
	g.Go(func() error {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", constant.Config.Server.Grpc.Port))
		if err != nil {
			return err
		}
		log.Printf(" [*] Grpc serve Waiting for connection:%d", constant.Config.Server.Grpc.Port)

		err = server.Grpc.Serve(lis)
		if err != nil {
			return err
		}
		return nil
	})
	err = g.Wait()
	if err != nil {
		panic(err)
	}
}
