package command

import (
	"github.com/go-juno/juno/internal/endpoint"
	"github.com/go-juno/juno/pkg/cli"
)

type CreateProjectCommand struct {
	endpoint *endpoint.Endpoints
}

func (t *CreateProjectCommand) Main() {}

func NewCreateProjectCommand(endpoint *endpoint.Endpoints) *cli.Command {
	return &cli.Command{
		Name:  "new",
		Short: "\tCreate a project\n",
		RunI: &CreateProjectCommand{
			endpoint: endpoint,
		},
	}
}
