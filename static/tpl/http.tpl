package http

import (
	"fmt"
	"net/http"
	"{{.Mod}}/api/http/middleware"
	"{{.Mod}}/internal/constant"
	"{{.Mod}}/internal/endpoint"
	"{{.Mod}}/static"

	"github.com/gin-gonic/gin"
)

func NewHttp(endpoints *endpoint.Endpoints) *http.Server {
	api := gin.New()
	api.Use(middleware.ErrMiddleware, middleware.LoggerFormate(), middleware.AuthMiddleware())
    {{- range .Funcs }}
    // {{.Description}}
    api.{{.Method}}("{{.Path}}", res.EndpointFunc(endpoints.{{.Name}}))
    {{- end }}


	Router(api, endpoints)
	s := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", constant.Config.Server.Host, constant.Config.Server.Http.Port),
		Handler: api,
	}
	return s
}
