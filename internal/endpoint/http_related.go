package endpoint

import (
	"github.com/go-juno/juno/pkg/util"
	"golang.org/x/xerrors"
)

type CreateHttpRequest struct {
	Name string
}

func (e *Endpoints) CreateHttpEndpoint(request *CreateHttpRequest) (err error) {
	//  获取mod
	mod, err := e.mod.GetMod()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	// 创建schema
	err = e.httpRelated.CreateScheme(mod, request.Name)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	// 创建serialize
	err = e.httpRelated.CreateSerialize(mod, request.Name)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	// 创建handle
	err = e.httpRelated.CreateHandle(mod, request.Name)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}

	// 更新http
	err = e.httpRelated.WireHttp(mod, request.Name)
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
