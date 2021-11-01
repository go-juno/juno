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
	CreateSchema(mod, name string, methodList []*Method) (err error)
	CreateSerialize(mod, name string, methodList []*Method) (err error)
	WireHttp(mod, name string) (err error)
	CreateHandle(mod, name string, methodList []*Method) (err error)
}

type httpRelatedService struct {
}

func genSchemaFile(mod, content string, methodList []*Method) (fileContent string) {

	if content == "" {
		fileContent = fmt.Sprintf(`package schema
import (
	"%s/internal/endpoint"
	"%s/internal/model"
)
			
`, mod, mod)
	} else {
		fileContent = content
	}

	for _, method := range methodList {
		if strings.Contains(fileContent, fmt.Sprintf("type %s struct", method.Request.Name)) {
			continue
		}

		var fieldString, transformFieldString, transformFunc string
		for _, field := range method.Request.FiledList {
			fieldString += fmt.Sprintf("\t%s %s %s\n", field.Name, field.Type, field.Tag)
			transformFieldString += fmt.Sprintf("\t\t%s: s.%s,\n", field.Name, field.Name)
		}

		if method.Request.IsList {
			transformFunc = fmt.Sprintf(`func (sList []*%s) Transform() []*endpoint.%sRequest {
				reqList := make([]*endpoint.%sRequest,len(s))
				for index,s :=range sList{
					reqList[index] =  &endpoint.%sRequest{
%s
					}
				}
				return reqList
			}
			`, method.Request.Name, method.Request.Name, method.Request.Name, method.Request.Name, transformFieldString)
		} else {
			transformFunc = fmt.Sprintf(`func (s *%s) Transform() *endpoint.%sRequest {
				req := &endpoint.%sRequest{
				%s
				}
				return req
			}
			`, method.Request.Name, method.Request.Name, method.Request.Name, transformFieldString)

		}

		structString := fmt.Sprintf(`type %s struct {
%s
}

%s
`, method.Request.Name, fieldString, transformFunc)
		fileContent += structString
	}

	return

}

func genSerializeFile(mod, content string, methodList []*Method) (fileContent string) {

	if content == "" {
		fileContent = fmt.Sprintf(`package serialize
import (
	"%s/internal/endpoint"
	"%s/internal/model"

	model2 "git.yupaopao.com/ops-public/kit/model"
)
	
`, mod, mod)

	} else {
		fileContent = content
	}

	for _, method := range methodList {
		if strings.Contains(fileContent, fmt.Sprintf("type %s struct", method.Response.Name)) {
			continue
		}
		var fieldString, transformFieldString, transformFunc string
		for _, field := range method.Response.FiledList {
			fieldString += fmt.Sprintf("\t%s %s %s\n", field.Name, field.Type, field.Tag)
			transformFieldString += fmt.Sprintf("\t\t%s: s.%s,\n", field.Name, field.Name)
		}

		if method.Response.Pagination {
			transformFunc = fmt.Sprintf(`func %sTransform(e *endpoint.%sResponse) (res *List) {
				items := make([]*%s, len(e.Items))
				res = &List{
					Total: e.Total,
				}
				for index, s := range e.Items {
					items[index] = &%s{
						%s
					}
				}
				res.Items = items
				return
			}
				`, method.Response.Name, method.Response.Name, method.Response.Name, method.Response.Name, transformFieldString)

		} else {
			if method.Response.IsList {
				transformFunc = fmt.Sprintf(`func %sTransform(e []*endpoint.%sResponse) (res []*%s) {
				res = make([]*%s, len(e))
				for index, s := range e {
					res[index] = &%s{
						%s
					}
				}
				return
			}
				`, method.Request.Name, method.Response.Name, method.Response.Name, method.Response.Name, method.Response.Name, transformFieldString)
			} else {
				transformFunc = fmt.Sprintf(`func %sTransform(s *endpoint.%sResponse) (res *%s) {
					if s != nil{
						res = &%s{
							%s
						}
					}
					return
				}
				`, method.Response.Name, method.Response.Name, method.Response.Name, method.Response.Name, transformFieldString)

			}
		}

		structString := fmt.Sprintf(`type %s struct {
%s
}

%s
`, method.Response.Name, fieldString, transformFunc)
		fileContent += structString
	}

	return

}

func (s *httpRelatedService) CreateSchema(mod, name string, methodList []*Method) (err error) {

	_, _, snake, _ := util.TransformName(name)
	httpSchemaDirPath := filepath.Join(util.GetPwd(), constant.HttpDirPath, "schema")
	schemaFileName := filepath.Join(httpSchemaDirPath, fmt.Sprintf("%s.go", snake))

	var ok bool
	ok, err = util.IsExistsFile(schemaFileName)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}

	var content, tpl string

	if ok {
		var contentByte []byte
		contentByte, err = util.ReadAll(schemaFileName)
		if err != nil {
			err = xerrors.Errorf("%w", err)
			return
		}
		content = string(contentByte)
	}

	tpl = genSchemaFile(mod, content, methodList)

	err = util.Mkdir(httpSchemaDirPath)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	err = util.WriteToFile(schemaFileName, tpl)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	return
}

func (s *httpRelatedService) CreateSerialize(mod, name string, methodList []*Method) (err error) {
	_, _, snake, _ := util.TransformName(name)
	httpSerializeDirPath := filepath.Join(util.GetPwd(), constant.HttpDirPath, "serialize")
	serializeFileName := filepath.Join(httpSerializeDirPath, fmt.Sprintf("%s.go", snake))

	var ok bool
	ok, err = util.IsExistsFile(serializeFileName)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	var tpl, content string
	if ok {
		var contentByte []byte
		contentByte, err = util.ReadAll(serializeFileName)
		if err != nil {
			err = xerrors.Errorf("%w", err)
			return
		}
		content = string(contentByte)
	}
	tpl = genSerializeFile(mod, content, methodList)

	err = util.Mkdir(httpSerializeDirPath)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	err = util.WriteToFile(serializeFileName, tpl)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	return
}

func (s *httpRelatedService) CreateHandle(mod, name string, methodList []*Method) (err error) {
	camel, class, snake, hyphen := util.TransformName(name)
	httpHandleDirPath := filepath.Join(util.GetPwd(), constant.HttpDirPath, "handle")
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
	httpDirPath := filepath.Join(util.GetPwd(), constant.HttpDirPath)
	err = util.Mkdir(httpDirPath)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	// wire add http
	httpFilePath := filepath.Join(httpDirPath, "http.go")
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
