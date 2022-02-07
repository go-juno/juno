package command

import (
	"fmt"

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

}

func NewCreateServiceCommand() *cli.Command {
	return &cli.Command{
		Name:  "service",
		Short: "\tCreate a service\n",
		RunI:  &CreateServiceCommand{},
	}
}
