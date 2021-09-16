package service

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewModService,
	NewBaseEnvService,
	NewProjectRelatedService,
	NewServiceRelatedService,
	NewEndpointRelatedService,
	NewHttpRelatedService,
	NewGrpcRelatedService,
)
