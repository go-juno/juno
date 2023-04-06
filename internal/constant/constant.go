package constant

var (
	TplMod          = []byte("{mod}")
	TplCamel        = []byte("{camel}")
	TplClass        = []byte("{class}")
	TplSnake        = []byte("{snake}")
	TplHyphen       = []byte("{hyphen}")
	Version         = "2.0.0"
	ServiceDirPath  = "internal/service"
	EndpointDirPath = "internal/endpoint"
	HttpDirPath     = "api/http"
	GrpcDirPath     = "api/grpc"
	WireFilePath    = "cmd/wire.go"
	ConstantPath    = "internal/constant"
	ModelPath       = "internal/entity"
)
