// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package cmd

import (
	"github.com/go-juno/juno/api/command"
	"github.com/go-juno/juno/internal/endpoint"
	"github.com/go-juno/juno/internal/service"
	"github.com/go-juno/juno/pkg/cli"
)

// Injectors from wire.go:

// initApp init endpoints
func InitServer() ([]*cli.Command, error) {
	baseEnvService := service.NewBaseEnvService()
	modService := service.NewModService()
	projectRelatedService := service.NewProjectRelatedService()
	serviceRelatedService := service.NewServiceRelatedService()
	endpointRelatedService := service.NewEndpointRelatedService()
	httpRelatedService := service.NewHttpRelatedService()
	grpcRelatedService := service.NewGrpcRelatedService()
	endpoints := endpoint.NewEndpoints(baseEnvService, modService, projectRelatedService, serviceRelatedService, endpointRelatedService, httpRelatedService, grpcRelatedService)
	v := command.GenCommandList(endpoints)
	return v, nil
}
