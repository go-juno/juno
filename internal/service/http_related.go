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

type HttpRelatedService interface {
	CreateScheme(mod, name string) (err error)
	CreateSerialize(mod, name string) (err error)
	WireHttp(mod, name string) (err error)
	CreateHandle(mod, name string) (err error)
}

type httpRelatedService struct {
}

func (s *httpRelatedService) CreateScheme(mod, name string) (err error) {

	camel, class, snake, hyphen := util.TransformName(name)
	httpSchemeDirPath := filepath.Join(constant.HttpDirPath, "schema")
	schemaFileName := filepath.Join(httpSchemeDirPath, fmt.Sprintf("%s.go", snake))

	var ok bool
	ok, err = util.IsExistsFile(schemaFileName)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	if ok {
		// TODO 根据更新endpoint 更新service
		return
	}

	err = util.Mkdir(httpSchemeDirPath)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	// 替换 模块
	tpl := util.Replace(static.HttpschemaTpl, mod, camel, class, snake, hyphen)
	err = util.WriteToFile(schemaFileName, tpl)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	return
}

func (s *httpRelatedService) CreateSerialize(mod, name string) (err error) {
	camel, class, snake, hyphen := util.TransformName(name)
	httpSerializeDirPath := filepath.Join(constant.HttpDirPath, "serialize")
	serializeFileName := filepath.Join(httpSerializeDirPath, fmt.Sprintf("%s.go", snake))

	var ok bool
	ok, err = util.IsExistsFile(serializeFileName)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	if ok {
		// TODO 根据更新endpoint 更新service
		return
	}

	err = util.Mkdir(httpSerializeDirPath)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	// 替换 模块
	tpl := util.Replace(static.HttpSerializeTpl, mod, camel, class, snake, hyphen)
	err = util.WriteToFile(serializeFileName, tpl)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	return
}

func (s *httpRelatedService) CreateHandle(mod, name string) (err error) {
	camel, class, snake, hyphen := util.TransformName(name)
	httpHandleDirPath := filepath.Join(constant.HttpDirPath, "handle")
	handleFileName := filepath.Join(httpHandleDirPath, fmt.Sprintf("%s.go", snake))

	var ok bool
	ok, err = util.IsExistsFile(handleFileName)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	if ok {
		// TODO 根据更新endpoint 更新service
		return
	}

	err = util.Mkdir(httpHandleDirPath)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}

	var apiPrefix string
	apiPrefix, err = util.GetApiPrefix()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}

	// 替换 模块
	tpl := util.ReplaceHttp(static.HttpHandleTpl, mod, camel, class, snake, hyphen, apiPrefix)
	err = util.WriteToFile(handleFileName, tpl)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	return
}

func (s *httpRelatedService) WireHttp(mod, name string) (err error) {
	camel, class, _, _ := util.TransformName(name)
	err = util.Mkdir(constant.HttpDirPath)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	// wire add http
	httpFilePath := filepath.Join(constant.HttpDirPath, "http.go")
	var content string = fmt.Sprintf(`
	package http
	import (
		"%s/internal/service"
		"github.com/google/wire"
	)
	type Https struct {
	}
		
	func NewHttps(
	) *Https {
		return &Https{
		}
	}
	var ProviderSet = wire.NewSet(NewHttps)
	`, mod)
	var ok bool
	ok, err = util.IsExistsFile(httpFilePath)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	if ok {
		var httpFile *os.File
		httpFile, err = os.Open(httpFilePath)
		if err != nil {
			err = xerrors.Errorf("%w", err)
			return
		}
		var c []byte
		c, err = ioutil.ReadAll(httpFile)
		if err != nil {
			err = xerrors.Errorf("%w", err)
			return
		}
		content = string(c)

	}
	structString := fmt.Sprintf(`type Https struct {
	%s service.%sService`, camel, class)
	paramString := fmt.Sprintf(`func NewHttps(
	%s service.%sService,`, camel, class)
	classString := fmt.Sprintf(`return &Https{
		%s:            %s,`, camel, camel)
	content = strings.ReplaceAll(content, "type Https struct {", structString)
	content = strings.ReplaceAll(content, "func NewHttps(", paramString)
	content = strings.ReplaceAll(content, "return &Https{", classString)
	err = util.WriteToFile(httpFilePath, content)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}

	return
}

func NewHttpRelatedService() HttpRelatedService {
	return &httpRelatedService{}
}
