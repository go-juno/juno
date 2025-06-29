package endpoint

import (
	"github.com/go-juno/juno/internal/service"
)

type CreateServiceRequest struct {
	Name string
	Kind string // Add Kind field for service type
}

func CreateServiceEndpoint(request *CreateServiceRequest) (err error) {
	return service.GenerateServiceRelatedFiles(request.Name, request.Kind)
}
