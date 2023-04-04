package service


import (
	"github.com/google/wire"
)


var ProviderSet = wire.NewSet(
	{{- range .ServiceList }}
 	{{ .NewFuncName }},
	{{- end }}
)
