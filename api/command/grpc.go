package command

import (
	"fmt"

	"github.com/go-juno/juno/pkg/cli"
	"github.com/go-juno/juno/pkg/cli/flag"
)

type CreateGrpcCommand struct {
}

func (t *CreateGrpcCommand) Main() {
	name := flag.Arguments().First().String()
	if name == "" {
		fmt.Println("grpc name required")
		return
	}
}

func NewCreateGrpcCommand() *cli.Command {
	return &cli.Command{
		Name:  "grpc",
		Short: "\tCreate a grpc\n",
		RunI:  &CreateGrpcCommand{},
	}
}
