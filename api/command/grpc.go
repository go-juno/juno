package command

import (
	"fmt"
	"log"

	"github.com/go-juno/juno/internal/endpoint"
	"github.com/go-juno/juno/pkg/cli"
	"github.com/go-juno/juno/pkg/cli/flag"
	"github.com/go-juno/juno/pkg/util"
)

type CreateGrpcCommand struct {
	endpoint *endpoint.Endpoints
}

func (t *CreateGrpcCommand) Main() {
	name := flag.Arguments().First().String()
	if name == "" {
		fmt.Println("grpc name required")
		return
	}
	req := &endpoint.CreateGrpcRequest{
		Name: name,
	}
	err := t.endpoint.CreateGrpcEndpoint(req)
	if err != nil {
		log.Println(util.Unwrap(err).Error())
	}

}

func NewCreateGrpcCommand(endpoint *endpoint.Endpoints) *cli.Command {
	return &cli.Command{
		Name:  "grpc",
		Short: "\tCreate a grpc\n",
		RunI: &CreateGrpcCommand{
			endpoint: endpoint,
		},
	}
}
