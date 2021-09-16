package command

import (
	"fmt"
	"log"

	"github.com/go-juno/juno/internal/endpoint"
	"github.com/go-juno/juno/pkg/cli"
	"github.com/go-juno/juno/pkg/cli/flag"
	"github.com/go-juno/juno/pkg/util"
)

type CreateServiceCommand struct {
	endpoint *endpoint.Endpoints
}

func (t *CreateServiceCommand) Main() {
	name := flag.Arguments().First().String()
	if name == "" {
		fmt.Println("service name required")
		return
	}
	req := &endpoint.CreateServiceRequest{
		Name: name,
	}
	err := t.endpoint.CreateServiceEndpoint(req)
	if err != nil {
		log.Println(util.Unwrap(err).Error())
	}

}

func NewCreateServiceCommand(endpoint *endpoint.Endpoints) *cli.Command {
	return &cli.Command{
		Name:  "service",
		Short: "\tCreate a service\n",
		RunI: &CreateServiceCommand{
			endpoint: endpoint,
		},
	}
}
