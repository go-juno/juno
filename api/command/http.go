package command

import (
	"fmt"

	"github.com/go-juno/juno/pkg/cli"
	"github.com/go-juno/juno/pkg/cli/flag"
)

type CreateHttpCommand struct {
}

func (t *CreateHttpCommand) Main() {
	name := flag.Arguments().First().String()
	if name == "" {
		fmt.Println("http name required")
		return
	}

}

func NewCreateHttpCommand() *cli.Command {
	return &cli.Command{
		Name:  "http",
		Short: "\tCreate a http\n",
		RunI:  &CreateHttpCommand{},
	}
}
