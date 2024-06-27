package http

import (
	"{{.}}/internal/endpoint"
	"{{.}}/api/http/middleware"
	"{{.}}/init/log"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewGin(endpoints *endpoint.Endpoints) (api *gin.Engine) {
	api = gin.New()
	gin.DefaultWriter = log.MultiWriter
	api.Use(middleware.ErrMiddleware, middleware.LoggerFormate(), middleware.AuthMiddleware())
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return api
}
