package endpoint

import "fmt"

type CreateGrpcRequest struct {
	Name string
}

func CreateGrpcEndpoint(req *CreateGrpcRequest) error {
	fmt.Printf("Endpoint: Creating grpc '%s'\n", req.Name)
	// This would call a grpc service, similar to the project service
	return nil
}
