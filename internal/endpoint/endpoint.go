package endpoint

import (
	"github.com/go-juno/juno/internal/service"
	"github.com/google/wire"
)

type Endpoints struct {
	file service.FileService
	mod  service.ModService
}

func NewEndpoints(
	file service.FileService,
	mod service.ModService,
) *Endpoints {
	return &Endpoints{
		file: file,
		mod:  mod,
	}
}

var ProviderSet = wire.NewSet(NewEndpoints)
