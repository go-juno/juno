package endpoint

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/go-juno/juno/internal/constant"
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

	srcDir := fmt.Sprintf("%s/pkg/mod/github.com/go-juno/juno@%s/example/juno", goPath, constant.Version)

	// 先安装依赖
	err = e.baseEnv.InstallEnv()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	log.Println("- Env Installation complete")

	// 复制脚手架
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	dest := fmt.Sprintf("%s/%s", pwd, request.Name)
	var ok bool
	ok, err = e.file.CopyPath(srcDir, dest)
	if !ok {
		panic(errors.New("Copy dir failed"))
	}
	log.Println(" > ok")
	if err := e.file.CreateMod(dest); err != nil {
		panic(errors.New("Replace go.mod failed"))
	}
	log.Println(" - Processing package name")
	if err := e.file.ReplaceAll(dest, "juno", request.Name); err != nil {
		panic(errors.New("Replace failed"))
	}
	log.Println(" > ok")
	log.Printf("Project '%s' is generated", request.Name)
	return
}
