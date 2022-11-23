package static

import (
	"embed"
)

//go:embed tpl/service.tpl
var ServiceTpl embed.FS

//go:embed tpl/mod.tpl
var ModTpl string

//go:embed tpl/service_wire.tpl
var ServiceWire embed.FS
