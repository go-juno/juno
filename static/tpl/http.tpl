package http

import (
	"fmt"
	"net/http"

	"{{.Mod}}/api/http/handle"
	"{{.Mod}}/api/http/middleware"
	"{{.Mod}}/internal/constant"
	"{{.Mod}}/internal/endpoint"
	"{{.Mod}}/static"

	"github.com/gin-gonic/gin"
)

func NewHttp(endpoints *endpoint.Endpoints) *http.Server {
	api := gin.New()
	api.Use(middleware.ErrMiddleware, middleware.LoggerFormate())
    {{- range .Funcs }}
    // {{.Description}}
    api.{{.Method}}("{{.Path}}", res.EndpointFunc(endpoints.{{.Name}}))
    {{- end }}

    api.GET("api/static/doc", func(ctx *gin.Context) {
		ctx.Header("Content-Type", "text/html; charset=utf-8")
		ctx.String(200, static.ApiDoc)
	})


	s := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", constant.Config.Server.Host, constant.Config.Server.Http.Port),
		Handler: api,
	}
	return s
}
