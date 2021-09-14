package command

import (
	"fmt"
	"log"

	"github.com/go-juno/juno/internal/endpoint"
	"github.com/go-juno/juno/pkg/cli"
	"github.com/go-juno/juno/pkg/cli/flag"
	"github.com/go-juno/juno/pkg/util"
)

type CreateProjectCommand struct {
	endpoint *endpoint.Endpoints
}

func (t *CreateProjectCommand) Main() {
	name := flag.Arguments().First().String()
	if name == "" {
		fmt.Println("project name required")
		return
	}
	req := &endpoint.CreateProjectRequest{
		Name: name,
	}
	err := t.endpoint.CreateProjectEndpoint(req)
	if err != nil {
		log.Println(util.Unwrap(err).Error())
	}

}

func NewCreateProjectCommand(endpoint *endpoint.Endpoints) *cli.Command {
	return &cli.Command{
		Name:  "new",
		Short: "\tCreate a project\n",
		RunI: &CreateProjectCommand{
			endpoint: endpoint,
		},
	}
}
