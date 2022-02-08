package static

import (
	_ "embed"
)

//go:embed tpl/service.tpl
var ServiceTpl string

//go:embed tpl/mod.tpl
var ModTpl string
