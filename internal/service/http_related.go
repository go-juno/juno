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
	_, class, _, _ := util.TransformName(name)
	err = util.Mkdir(constant.HttpDirPath)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	// wire add http
	httpFilePath := filepath.Join(constant.HttpDirPath, "http.go")
	var content string = fmt.Sprintf(`package http
	import (
		"fmt"
		"%s/api/http/middleware"
		"net/http"
	
		"%s/api/http/handle"
		"%s/internal/constant"
		"%s/internal/endpoint"
		"%s/static"
	
		"github.com/gin-gonic/gin"
	)
	
	func NewHttp(endpoints *endpoint.Endpoints) *http.Server {
		ginEngine := gin.New()
		ginEngine.Use(middleware.ErrMiddleware)
		api := ginEngine.Group("/api")
	    api.GET("/doc", func(ctx *gin.Context) {
		ctx.Header("Content-Type", "text/html; charset=utf-8")
		ctx.String(200, static.ApiDoc)
	})
	s := &http.Server{
		Addr:    fmt.Sprintf(":%%d", constant.Config.Server.Http.Port),
		Handler: ginEngine,
	}
	return s
}`, mod, mod, mod, mod, mod)
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
		exitsService := fmt.Sprintf(`%sBluePrint`, class)
		if strings.Contains(content, exitsService) {
			return
		}

	}
	structString := fmt.Sprintf(`handle.%sBluePrint(api, endpoints)
	s := &http.Server{`, class)
	content = strings.ReplaceAll(content, "s := &http.Server{", structString)
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
