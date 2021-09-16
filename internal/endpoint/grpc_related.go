package endpoint

import (
	"github.com/go-juno/juno/pkg/util"
	"golang.org/x/xerrors"
)

type CreateGrpcRequest struct {
	Name string
}

func (e *Endpoints) CreateGrpcEndpoint(request *CreateGrpcRequest) (err error) {
	//  获取mod
	mod, err := e.mod.GetMod()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}

	// 创建proto
	err = e.grpcRelated.CreateProto(mod, request.Name)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}

	// 创建Service
	err = e.grpcRelated.CreateService(mod, request.Name)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}

	// Wire Grpc
	err = e.grpcRelated.WireGrpc(mod, request.Name)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}

	err = util.FmtCode()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	return
}
