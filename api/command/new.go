package command

import (
	"github.com/go-juno/juno/internal/endpoint"
	"github.com/go-juno/juno/pkg/cli"
)

type NewCommand struct {
	endpoint *endpoint.Endpoints
}

func (t *NewCommand) Main() {}

func NewNewCommand(endpoint *endpoint.Endpoints) *cli.Command {
	return &cli.Command{
		Name:  "new",
		Short: "\tCreate a project\n",
		RunI: &NewCommand{
			endpoint: endpoint,
		},
	}
}
