package constant

import "github.com/go-juno/juno/init/version"

var (
	TplMod          = []byte("{mod}")
	TplCamel        = []byte("{camel}")
	TplClass        = []byte("{class}")
	TplSnake        = []byte("{snake}")
	TplHyphen       = []byte("{hyphen}")
	Version         = version.Version
	ServiceDirPath  = "internal/service"
	EndpointDirPath = "internal/endpoint"
	HttpDirPath     = "api/http"
	GrpcDirPath     = "api/grpc"
	WireFilePath    = "cmd/wire.go"
	ContantPath     = "internal/contant"
	ModelPath       = "internal/model"
)
