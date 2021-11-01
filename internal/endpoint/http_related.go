package endpoint

import (
	"fmt"
	"path/filepath"

	"github.com/go-juno/juno/internal/constant"
	"github.com/go-juno/juno/internal/service"
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
	// 读取endpoint内容
	_, _, snake, _ := util.TransformName(request.Name)
	fileName := filepath.Join(util.GetPwd(), constant.EndpointDirPath, fmt.Sprintf("%s.go", snake))
	var contentByte []byte
	contentByte, err = util.ReadAll(fileName)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}

	//  解析endpoint
	var methodList []*service.Method
	methodList, err = e.parse.GenMethodList(string(contentByte))
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}

	// 创建schema
	err = e.httpRelated.CreateSchema(mod, request.Name, methodList)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	// 创建serialize
	err = e.httpRelated.CreateSerialize(mod, request.Name, methodList)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	// 创建handle
	err = e.httpRelated.CreateHandle(mod, request.Name, methodList)
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
