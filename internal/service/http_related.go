package service

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-juno/juno/internal/constant"
	"github.com/go-juno/juno/pkg/util"
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

func genHandleFile(mod, apiPrefix, fileCamel, fileHyphen, fileClass, content string, methodList []*Method) (fileContent string) {

	if content == "" {
		fileContent = fmt.Sprintf(`package handle
import (
		"%s/api/http/schema"
		"%s/api/http/serialize"
		"%s/internal/endpoint"
		"%s/pkg/res"
	
		"github.com/gin-gonic/gin"
		"golang.org/x/xerrors"
)	
`, mod, mod, mod, mod)

	} else {
		fileContent = content
	}

	var bluePrintString, methodString string

	for _, method := range methodList {
		camel, class, _, _ := util.TransformName(method.Name)
		if strings.Contains(fileContent, fmt.Sprintf("%s(endpoints)", camel)) {
			continue
		}
		if bluePrintString != "" {
			bluePrintString += "\n"
		}
		bluePrintString += fmt.Sprintf(`r.%s("", %s(endpoints))`, util.GetMethod(camel), camel)
		var requestString, responseString, successResponseString string
		for _, field := range method.Request.FiledList {
			if requestString != "" {
				requestString += "\n"
			}
			_, _, s, _ := util.TransformName(field.Name)
			requestString += fmt.Sprintf("@apiParam {%s} %s %s", field.Type, s, s)

		}
		for _, field := range method.Response.FiledList {
			if responseString != "" {
				responseString += "\n"
			}

			_, _, s, _ := util.TransformName(field.Name)
			responseString += fmt.Sprintf("@apiSuccess {%s} %s %s", field.Type, s, s)
			var value string
			if strings.Contains(field.Type, "int") {
				value = "1"
			} else {
				value = fmt.Sprintf(`"%s"`, field.Name)
			}
			if successResponseString != "" {
				successResponseString += ","
			}
			successResponseString += fmt.Sprintf(`"%s":%s`, s, value)
		}
		if method.Response.Pagination {
			successResponseString = fmt.Sprintf(`{"status":"success","msg":null,"data":{"items":[{%s}],"total":10}}`, successResponseString)
		} else if method.Response.IsList {
			successResponseString = fmt.Sprintf(`{"status":"success","msg":null,"data":[{%s}]}`, successResponseString)
		} else {
			successResponseString = fmt.Sprintf(`{"status":"success","msg":null,"data":{%s}}`, successResponseString)
		}
		methodString += fmt.Sprintf(`/**
@api {%s} /api%s/%s %s
@apiVersion 1.0.0
@apiName %s
@apiGroup %s
%s
@apiSuccess {string} status 状态码
@apiSuccess {string} msg 返回的信息
@apiSuccess {list} data 返回的数据
%s
@apiSuccessExample Success-Response:
%s
@apiErrorExample Error-Response:
{"status": "failure", "data": null, "message": "失败"}
**/
func %s(endpoints *endpoint.Endpoints) gin.HandlerFunc {
	return func(c *gin.Context) {
		var schema schema.%s
		err := c.ShouldBind(&schema)
		if err != nil {
			err = xerrors.Errorf("%%w", err)
			res.ParamCheckRes(c, err)
			return
		}
		req := schema.Transform()
		result, err := endpoints.%sEndpoint(c, req)
		if err != nil {
			err = xerrors.Errorf("%%w", err)
			res.FailureRes(c, err)
			return
		}
		res.SuccessRes(c, serialize.%sTransform(result))
	}
}
`, util.GetMethod(camel), apiPrefix, fileHyphen, camel, camel, fileCamel, requestString, responseString, successResponseString, camel, class, class, class)

	}

	if bluePrintString != "" {
		fileContent = fmt.Sprintf(`%s
		func %sBluePrint(v1 *gin.RouterGroup, endpoints *endpoint.Endpoints) {
			r := v1.Group("/%s")
			%s
		}
		`, fileContent, fileClass, fileHyphen, bluePrintString)
	}
	if methodString != "" {
		fileContent = fmt.Sprintf(`%s
%s
		`, fileContent, methodString)
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
	var tpl, content string
	if ok {
		var contentByte []byte
		contentByte, err = util.ReadAll(handleFileName)
		if err != nil {
			err = xerrors.Errorf("%w", err)
			return
		}
		content = string(contentByte)
	}
	var apiPrefix string
	apiPrefix, err = util.GetApiPrefix()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	tpl = genHandleFile(mod, apiPrefix, camel, hyphen, class, content, methodList)

	err = util.Mkdir(httpHandleDirPath)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}

	// 写入文件
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
