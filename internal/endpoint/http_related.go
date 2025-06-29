package endpoint

import (
	"github.com/go-juno/juno/internal/service"
)

type CreateHttpRequest struct {
}

func CreateHttpEndpoint(request *CreateHttpRequest) (err error) {
	return service.GenerateHttpRelatedFiles()
}
