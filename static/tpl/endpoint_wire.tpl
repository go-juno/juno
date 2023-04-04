package endpoint

import (
    "{{ .Mod }}/internal/service"
	"github.com/google/wire"
)




type Endpoints struct { 
	{{- range .ServiceList }}
    {{ .Name.Camel }} service.{{ .ServiceName }}
    {{- end }}
}

func NewEndpoints(
	{{- range .ServiceList }}
    {{ .Name.Camel }} service.{{ .ServiceName }},
    {{- end }}
) *Endpoints {
	return &Endpoints{
	{{- range .ServiceList }}
    {{ .Name.Camel }}:{{ .Name.Camel }},
    {{- end }}
	}
}

var ProviderSet = wire.NewSet(NewEndpoints)
