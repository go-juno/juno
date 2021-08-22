package endpoint

import (
	"github.com/google/wire"
)

type Endpoints struct {
	new service.NewService
}

func NewEndpoints(
	new service.NewService) *Endpoints {
	return &Endpoints{
		new: new}
}

var ProviderSet = wire.NewSet(NewEndpoints)
