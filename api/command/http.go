package command

import (
	"log"

	"github.com/go-juno/juno/internal/endpoint"
	"github.com/go-juno/juno/pkg/cli"
)

type CreateHttpCommand struct {
}

func (t *CreateHttpCommand) Main() {
	err := endpoint.CreateHttpEndpoint(&endpoint.CreateHttpRequest{})
	if err != nil {
		log.Printf("err: %+v", err)
	}
}

func NewCreateHttpCommand() *cli.Command {
	return &cli.Command{
		Name:  "http",
		Short: "\tCreate a http\n",
		RunI:  &CreateHttpCommand{},
	}
}
