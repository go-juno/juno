// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package cmd

import (
	"github.com/go-juno/juno/api/command"
	"github.com/go-juno/juno/internal/endpoint"
	"github.com/go-juno/juno/internal/service"
	"github.com/go-juno/juno/pkg/cli"
	"github.com/google/wire"
)

// initApp init endpoints
func InitServer() ([]*cli.Command, error) {
	panic(wire.Build(
		service.ProviderSet,
		endpoint.ProviderSet,
		command.GenCommandList,
	))
}
