package static

import (
	_ "embed"
)

//go:embed tpl/service.tpl
var ServiceTpl string

//go:embed tpl/endpoint.tpl
var EndpointTpl string

//go:embed tpl/http_schema.tpl
var HttpschemaTpl string

//go:embed tpl/http_serialize.tpl
var HttpSerializeTpl string

//go:embed tpl/http_handle.tpl
var HttpHandleTpl string

//go:embed tpl/grpc_service.tpl
var GrpcService string

//go:embed tpl/grpc_proto.tpl
var GrpcProto string

//go:embed tpl/mod.tpl
var ModTpl string
