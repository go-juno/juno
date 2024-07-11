package http

import (
	"fmt"
	"net/http"
	"{{.Mod}}/api/http/middleware"
	"{{.Mod}}/internal/constant"
	"{{.Mod}}/internal/endpoint"
	"{{.Mod}}/pkg/res"
	
	"github.com/gin-gonic/gin"
)

func NewHttp(api *gin.Engine, endpoints *endpoint.Endpoints) *http.Server{
    {{- range .Funcs }}
    api.{{.Method}}("{{.Path}}", res.EndpointFunc(endpoints.{{.Name}}))
    {{- end }}

	s := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", constant.Config.Server.Host, constant.Config.Server.Http.Port),
		Handler: api,
	}
	return s
}
