package endpoint

import (
	"github.com/go-juno/juno/pkg/util"
	"golang.org/x/xerrors"
)

type CreateProjectRequest struct {
	Name string
}

func (e *Endpoints) CreateProjectEndpoint(request *CreateProjectRequest) (err error) {
	//  检查包地址
	var goPath string
	goPath, err = e.baseEnv.GetGoEnvPath()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	// 先安装依赖
	err = e.baseEnv.InstallEnv()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	// 生成文件
	err = e.projectRelated.CreateProject(request.Name, goPath)
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
