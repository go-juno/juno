package endpoint

import (
	"github.com/go-juno/juno/internal/service"
	"github.com/google/wire"
)

type Endpoints struct {
	file service.FileService
	mod  service.ModService
	new  service.NewService
}

func NewEndpoints(
	file service.FileService,
	mod service.ModService,
	new service.NewService) *Endpoints {
	return &Endpoints{
		file: file,
		mod:  mod,
		new:  new,
	}
}

var ProviderSet = wire.NewSet(NewEndpoints)
