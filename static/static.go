package static

import (
	_ "embed"
)

//go:embed tpl/service.tpl
var ServiceTpl string

//go:embed tpl/mod.tpl
var ModTpl string

//go:embed tpl/service_wire.tpl
var ServiceWireTpl string

//go:embed tpl/endpoint.tpl
var EndpointTpl string

//go:embed tpl/endpoint_wire.tpl
var EndpointWireTpl string

//go:embed tpl/http.tpl
var HttpTpl string

//go:embed tpl/handle.tpl
var HandleTpl string

//go:embed tpl/mongo_model.tpl
var MongoModelTpl string

//go:embed tpl/mongo_service.tpl
var MongoServiceTpl string

//go:embed tpl/main.tpl
var MainTpl string

//go:embed tpl/readme.tpl
var ReadmeTpl string

//go:embed tpl/gitignore.tpl
var GitignoreTpl string

//go:embed tpl/dockerfile.tpl
var DockerfileTpl string

//go:embed tpl/air.tpl
var AirTpl string

//go:embed tpl/pre_commit.tpl
var PreCommitTpl string
