package endpoint

import (
	"github.com/go-juno/juno/internal/service"
	"github.com/google/wire"
)

type Endpoints struct {
	baseEnv         service.BaseEnvService
	mod             service.ModService
	projectRelated  service.ProjectRelatedService
	serviceRelated  service.ServiceRelatedService
	endpointRelated service.EndpointRelatedService
	httpRelated     service.HttpRelatedService
}

func NewEndpoints(
	baseEnv service.BaseEnvService,
	mod service.ModService,
	projectRelated service.ProjectRelatedService,
	serviceRelated service.ServiceRelatedService,
	endpointRelated service.EndpointRelatedService,
	httpRelated service.HttpRelatedService,
) *Endpoints {
	return &Endpoints{
		baseEnv:         baseEnv,
		mod:             mod,
		projectRelated:  projectRelated,
		serviceRelated:  serviceRelated,
		endpointRelated: endpointRelated,
		httpRelated:     httpRelated,
	}
}

var ProviderSet = wire.NewSet(NewEndpoints)
