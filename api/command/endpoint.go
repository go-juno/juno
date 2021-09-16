package command

import (
	"fmt"
	"log"

	"github.com/go-juno/juno/internal/endpoint"
	"github.com/go-juno/juno/pkg/cli"
	"github.com/go-juno/juno/pkg/cli/flag"
	"github.com/go-juno/juno/pkg/util"
)

type CreateEndpointCommand struct {
	endpoint *endpoint.Endpoints
}

func (t *CreateEndpointCommand) Main() {
	name := flag.Arguments().First().String()
	if name == "" {
		fmt.Println("endpoint name required")
		return
	}
	req := &endpoint.CreateEndpointRequest{
		Name: name,
	}
	err := t.endpoint.CreateEndpointEndpoint(req)
	if err != nil {
		log.Println(util.Unwrap(err).Error())
	}

}

func NewCreateEndpointCommand(endpoint *endpoint.Endpoints) *cli.Command {
	return &cli.Command{
		Name:  "endpoint",
		Short: "\tCreate a endpoint\n",
		RunI: &CreateEndpointCommand{
			endpoint: endpoint,
		},
	}
}
