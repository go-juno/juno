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

type EndpointRelatedService interface {
	CreateEndpoint(mod, name string) (err error)
	WireEndpoint(mod, name string) (err error)
}

type endpointRelatedService struct {
}

func (s *endpointRelatedService) CreateEndpoint(mod, name string) (err error) {

	camel, class, snake, hyphen := util.TransformName(name)
	fileName := filepath.Join(constant.EndpointDirPath, fmt.Sprintf("%s.go", snake))
	err = util.Mkdir(constant.EndpointDirPath)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}

	// 替换 模块
	tpl := util.Replace(static.EndpointTpl, mod, camel, class, snake, hyphen)
	err = util.WriteToFile(fileName, tpl)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	return
}

func (s *endpointRelatedService) WireEndpoint(mod, name string) (err error) {
	camel, class, _, _ := util.TransformName(name)
	err = util.Mkdir(constant.EndpointDirPath)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	// wire add endpoint
	endpointFilePath := filepath.Join(constant.EndpointDirPath, "endpoint.go")
	var content string = fmt.Sprintf(`
	package endpoint
	import (
		"%s/internal/service"
		"github.com/google/wire"
	)
	type Endpoints struct {
	}
		
	func NewEndpoints(
	) *Endpoints {
		return &Endpoints{
		}
	}
	var ProviderSet = wire.NewSet(NewEndpoints)
	`, mod)
	var ok bool
	ok, err = util.IsExistsFile(endpointFilePath)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	if ok {
		var endpointFile *os.File
		endpointFile, err = os.Open(endpointFilePath)
		if err != nil {
			err = xerrors.Errorf("%w", err)
			return
		}
		var c []byte
		c, err = ioutil.ReadAll(endpointFile)
		if err != nil {
			err = xerrors.Errorf("%w", err)
			return
		}
		content = string(c)

	}
	structString := fmt.Sprintf(`type Endpoints struct {
	%s service.%sService`, camel, class)
	paramString := fmt.Sprintf(`func NewEndpoints(
	%s service.%sService,`, camel, class)
	classString := fmt.Sprintf(`return &Endpoints{
		%s:            %s,`, camel, camel)
	content = strings.ReplaceAll(content, "type Endpoints struct {", structString)
	content = strings.ReplaceAll(content, "func NewEndpoints(", paramString)
	content = strings.ReplaceAll(content, "return &Endpoints{", classString)
	err = util.WriteToFile(endpointFilePath, content)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}

	return
}

func NewEndpointRelatedService() EndpointRelatedService {
	return &endpointRelatedService{}
}
