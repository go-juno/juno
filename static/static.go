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
var EndpointWire string
