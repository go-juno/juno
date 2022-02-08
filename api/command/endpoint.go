package command

import (
	"fmt"

	"github.com/go-juno/juno/pkg/cli"
	"github.com/go-juno/juno/pkg/cli/flag"
)

type CreateEndpointCommand struct {
}

func (t *CreateEndpointCommand) Main() {
	name := flag.Arguments().First().String()
	if name == "" {
		fmt.Println("endpoint name required")
		return
	}
	// req := &endpoint.CreateEndpointRequest{
	// 	Name: name,
	// }
	// err := t.endpoint.CreateEndpointEndpoint(req)
	// if err != nil {
	// 	log.Printf("err: %+v", err)
	// }

}

func NewCreateEndpointCommand() *cli.Command {
	return &cli.Command{
		Name:  "endpoint",
		Short: "Create a endpoint\n",
		RunI:  &CreateEndpointCommand{},
	}
}
