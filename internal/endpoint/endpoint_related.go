package endpoint

import (
	"github.com/go-juno/juno/internal/service"
	"github.com/go-juno/juno/pkg/command"
	"github.com/go-juno/juno/pkg/util"
	"golang.org/x/xerrors"
)

type CreateEndpointRequest struct {
	Name string
}

func CreateEndpointEndpoint(request *CreateEndpointRequest) (err error) {
	//  获取mod
	mod, err := util.GetMod()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	// 先创建内容
	err = service.GeneratorEndpoint(mod, request.Name)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	// import信息更新
	err = command.GoimportsCode()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	return
}
