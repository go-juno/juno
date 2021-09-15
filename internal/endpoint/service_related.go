package endpoint

import (
	"github.com/go-juno/juno/pkg/util"
	"golang.org/x/xerrors"
)

type CreateServiceRequest struct {
	Name string
}

func (e *Endpoints) CreateServiceEndpoint(request *CreateServiceRequest) (err error) {
	//  获取mod
	mod, err := e.mod.GetMod()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	// 先创建内容
	err = e.serviceRelated.CreateService(mod, request.Name)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	// 更新servcie wire
	err = e.serviceRelated.WireService(request.Name)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}

	// 更新endpoint wire
	err = e.endpointRelated.WireEndpoint(mod, request.Name)
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
