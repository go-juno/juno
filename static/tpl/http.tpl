package http

import (
	"fmt"
	"net/http"
	"{{.Mod}}/api/http/middleware"
	"{{.Mod}}/internal/constant"
	"{{.Mod}}/internal/endpoint"
	"{{.Mod}}/static"
	_ "ep-service/docs"
	
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewHttp(endpoints *endpoint.Endpoints) *http.Server {
	api := gin.New()
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	api.Use(middleware.ErrMiddleware, middleware.LoggerFormate(), middleware.AuthMiddleware())
    {{- range .Funcs }}
    api.{{.Method}}("{{.Path}}", res.EndpointFunc(endpoints.{{.Name}}))
    {{- end }}


	Router(api, endpoints)
	s := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", constant.Config.Server.Host, constant.Config.Server.Http.Port),
		Handler: api,
	}
	return s
}
