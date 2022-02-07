package command

import (
	"github.com/go-juno/juno/pkg/cli"
)

func GenCommandList() (cmdList []*cli.Command) {
	cmdList = append(cmdList,
		NewCreateProjectCommand(),
		NewCreateServiceCommand(),
		NewCreateEndpointCommand(),
		NewCreateHttpCommand(),
		NewCreateGrpcCommand(),
	)
	return
}
