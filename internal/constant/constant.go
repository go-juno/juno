package constant

import "github.com/go-juno/juno/init/version"

var (
	TplMod          = "{mod}"
	TplCamel        = "{camel}"
	TplClass        = "{class}"
	TplSnake        = "{snake}"
	TplHyphen       = "{hyphen}"
	Version         = version.Version
	ServiceDirPath  = "internal/service"
	EndpointDirPath = "internal/endpoint"
	HttpDirPath     = "api/http"
	ReplaceMod      = "github.com/go-juno/juno/example/juno"
	GrpcDirPath     = "api/grpc"
	WireFilePath    = "cmd/wire.go"
)
