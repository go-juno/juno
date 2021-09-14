package http

import (
	"fmt"
	"net/http"

	"juno/api/http/handle"
	"juno/api/http/middleware"
	"juno/internal/constant"
	"juno/internal/endpoint"
	"juno/static"

	"github.com/gin-gonic/gin"
)

func NewHttp(endpoints *endpoint.Endpoints) *http.Server {
	ginEngine := gin.New()
	ginEngine.Use(middleware.ErrMiddleware)
	api := ginEngine.Group("/api")
	handle.GreetingBluePrint(api, endpoints)

	api.GET("/doc", func(ctx *gin.Context) {
		ctx.Header("Content-Type", "text/html; charset=utf-8")
		ctx.String(200, static.ApiDoc)
	})
	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", constant.Config.Server.Http.Port),
		Handler: ginEngine,
	}
	return s
}
