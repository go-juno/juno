package constant

import "github.com/go-juno/juno/init/version"

var (
	Version         = version.Version
	ServiceDirPath  = "internal/service"
	EndpointDirPath = "internal/endpoint"
	HttpDirPath     = "api/http"
	ReplaceMod      = "github.com/go-juno/juno/example/juno"
	TplMod          = "juno"
	TplCamel        = "greeting"
	TplClass        = "Greeting"
	TplSnake        = "greet_sb"
	TplHyphen       = "greeting-router"
	GrpcDirPath     = "api/grpc"
)
