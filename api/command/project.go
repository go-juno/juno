package command

import (
	"fmt"

	"github.com/go-juno/juno/pkg/cli"
	"github.com/go-juno/juno/pkg/cli/flag"
)

type CreateProjectCommand struct {
}

func (t *CreateProjectCommand) Main() {
	name := flag.Arguments().First().String()
	if name == "" {
		fmt.Println("project name required")
		return
	}

}

func NewCreateProjectCommand() *cli.Command {
	return &cli.Command{
		Name:  "new",
		Short: "\tCreate a project\n",
		RunI:  &CreateProjectCommand{},
	}
}
