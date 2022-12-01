package endpoint

import (
	"context"
	"golang.org/x/xerrors"
    {{ range .Packages }}
    {{ . }}
    {{- end }}
)


    {{ range .Funcs }}
type {{.Name}}Request struct {
    {{- range .Request }}
    {{- if (and (ne .TypeString "error") (ne .TypeString "context.Context"))}}
    {{ .Name }}  {{ .TypeString }}
    {{- end }}
    {{- end }}
}

type {{.Name}}Response struct {
    {{- range .Response }}
    {{- if (and (ne .TypeString "error") (ne .TypeString "context.Context"))}}
    {{ .Name }}  {{ .TypeString }}
    {{- end }}
    {{- end }}
}



    {{ end }}
