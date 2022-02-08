package command

import (
	"fmt"
	"log"

	"github.com/go-juno/juno/internal/endpoint"
	"github.com/go-juno/juno/pkg/cli"
	"github.com/go-juno/juno/pkg/cli/flag"
)

type CreateServiceCommand struct {
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
	err := endpoint.CreateServiceEndpoint(req)
	if err != nil {
		log.Printf("err: %+v", err)
	}

}

func NewCreateServiceCommand() *cli.Command {
	return &cli.Command{
		Name:  "service",
		Short: "\tCreate a service\n",
		RunI:  &CreateServiceCommand{},
	}
}
