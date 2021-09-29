package command

import (
	"fmt"
	"log"

	"github.com/go-juno/juno/internal/endpoint"
	"github.com/go-juno/juno/pkg/cli"
	"github.com/go-juno/juno/pkg/cli/flag"
)

type CreateHttpCommand struct {
	endpoint *endpoint.Endpoints
}

func (t *CreateHttpCommand) Main() {
	name := flag.Arguments().First().String()
	if name == "" {
		fmt.Println("http name required")
		return
	}
	req := &endpoint.CreateHttpRequest{
		Name: name,
	}
	err := t.endpoint.CreateHttpEndpoint(req)
	if err != nil {
		// log.Println(util.Unwrap(err).Error())
		log.Printf("err: %+v", err)
	}

}

func NewCreateHttpCommand(endpoint *endpoint.Endpoints) *cli.Command {
	return &cli.Command{
		Name:  "http",
		Short: "\tCreate a http\n",
		RunI: &CreateHttpCommand{
			endpoint: endpoint,
		},
	}
}
