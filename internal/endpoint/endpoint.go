package endpoint

import (
	"github.com/go-juno/juno/internal/service"
	"github.com/google/wire"
)

type Endpoints struct {
	baseEnv service.BaseEnvService
	file    service.FileService
	mod     service.ModService
}

func NewEndpoints(
	baseEnv service.BaseEnvService,
	file service.FileService,
	mod service.ModService,
) *Endpoints {
	return &Endpoints{
		baseEnv: baseEnv,
		file:    file,
		mod:     mod,
	}
}

var ProviderSet = wire.NewSet(NewEndpoints)
