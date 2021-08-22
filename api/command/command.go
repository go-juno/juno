package command

import (
	"github.com/go-juno/juno/internal/endpoint"
	"github.com/go-juno/juno/pkg/cli"
)

func GenCommandList(endpoint *endpoint.Endpoints) (cmdList []*cli.Command) {
	cmdList = append(cmdList, NewNewCommand(endpoint))
	return
}
