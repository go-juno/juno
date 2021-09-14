package endpoint

import (
	"juno/internal/service"

	"github.com/google/wire"
)

type Endpoints struct {
	greeting service.GreetingService
}

func NewEndpoints(
	greeting service.GreetingService,

) *Endpoints {
	return &Endpoints{
		greeting: greeting,
	}
}

var ProviderSet = wire.NewSet(NewEndpoints)
