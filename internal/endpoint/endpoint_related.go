package endpoint

import (
	"github.com/go-juno/juno/pkg/util"
	"golang.org/x/xerrors"
)

type CreateEndpointRequest struct {
	Name string
}

func (e *Endpoints) CreateEndpointEndpoint(request *CreateEndpointRequest) (err error) {
	//  获取mod
	mod, err := e.mod.GetMod()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	// 先创建内容
	err = e.endpointRelated.CreateEndpoint(mod, request.Name)
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
