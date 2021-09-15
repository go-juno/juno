package service

import (
	"fmt"
	"log"
	"os"

	"github.com/go-juno/juno/internal/constant"
	"github.com/go-juno/juno/pkg/util"
	"golang.org/x/xerrors"
)

type ProjectRelatedService interface {
	CreateProject(name, goPath string) (err error)
}

type projectRelatedService struct {
}

func (s *projectRelatedService) CreateProject(name, goPath string) (err error) {

	srcDir := fmt.Sprintf("%s/pkg/mod/github.com/go-juno/juno@%s/example/juno", goPath, constant.Version)
	// 复制脚手架
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	dest := fmt.Sprintf("%s/%s", pwd, name)
	var ok bool
	ok, err = util.CopyPath(srcDir, dest)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	if !ok {
		err = xerrors.New("Copy dir failed")
		return
	}
	err = util.CreateMod(dest)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	err = util.ReplaceAll(dest, constant.ReplaceMod, name)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	log.Printf("Project '%s' is generated", name)
	return
}

func NewProjectRelatedService() ProjectRelatedService {
	return &projectRelatedService{}
}
