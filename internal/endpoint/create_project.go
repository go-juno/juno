package endpoint

import (
	"context"
	"log"

	"golang.org/x/xerrors"
)

type CreateProjectRequest struct {
}

func (e *Endpoints) CreateProjectEndpoint(ctx context.Context, request *CreateProjectRequest) (err error) {
	// 先安装依赖
	err = e.baseEnv.InstallEnv()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	log.Println("- Env Installation complete")
	//
	return
}
