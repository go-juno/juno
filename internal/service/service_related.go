package service

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-juno/juno/internal/constant"
	"github.com/go-juno/juno/pkg/util"
	"github.com/go-juno/juno/static"
	"golang.org/x/xerrors"
)

func GeneratorService(mod, name string) (err error) {
	camel, class, snake, hyphen := util.TransformName(name)
	serviceDirPath := filepath.Join(util.GetPwd(), constant.ServiceDirPath)
	fileName := filepath.Join(serviceDirPath, fmt.Sprintf("%s.go", snake))
	var ok bool
	ok, err = util.IsExistsFile(fileName)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	if ok {
		err = xerrors.Errorf("file:%s already exists", fileName)
		return
	}
	err = util.Mkdir(serviceDirPath)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}

	// 替换 模块
	tpl := util.Replace(static.ServiceTpl, mod, camel, class, snake, hyphen)
	err = util.WriteToFile(fileName, tpl)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	return
}

func WireService(name string) (err error) {
	_, class, _, _ := util.TransformName(name)
	// wire add service
	serviceDirPath := filepath.Join(util.GetPwd(), constant.ServiceDirPath)
	serviceFilePath := filepath.Join(serviceDirPath, "service.go")
	var content string = `
		package service
		import "github.com/google/wire"
		var ProviderSet = wire.NewSet(
		)
		`
	var ok bool
	ok, err = util.IsExistsFile(serviceFilePath)

	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	if ok {
		var serviceFile *os.File
		serviceFile, err = os.Open(serviceFilePath)
		if err != nil {
			err = xerrors.Errorf("%w", err)
			return
		}
		var c []byte
		c, err = ioutil.ReadAll(serviceFile)
		if err != nil {
			err = xerrors.Errorf("%w", err)
			return
		}
		content = string(c)
		exitsService := fmt.Sprintf(`New%sService`, class)
		if strings.Contains(content, exitsService) {
			return
		}

	}
	newServiceString := fmt.Sprintf(`New%sService,
	)`, class)
	file := strings.ReplaceAll(content, ")", newServiceString)
	err = util.WriteToFile(serviceFilePath, file)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	return
}
