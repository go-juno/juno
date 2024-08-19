package endpoint

import (
	"github.com/go-juno/juno/internal/service"
	"github.com/go-juno/juno/pkg/command"
	"github.com/go-juno/juno/pkg/util"
	"golang.org/x/xerrors"
)

type CreateHttpRequest struct {
}

func CreateHttpEndpoint(request *CreateHttpRequest) (err error) {
	//  获取mod
	mod, err := util.GetMod()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	// 更新servcie wire
	err = service.WireService(mod)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}

	// 更新endpoint wire
	err = service.WireEndpoint(mod)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	// 创建handle
	err = service.GenerateHandle(mod)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}

	// 先创建内容
	err = service.GenerateHttp(mod)
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

	err = command.RunWire()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	return
}
