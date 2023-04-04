package endpoint

import (
    {{ range .Packages }}
    {{ . }}
    {{- end }}
)


    {{ range .Funcs }}
type {{.Name.Class }}Request struct {
    {{- range .Request }}
    {{- if (and (ne .TypeString "error") (ne .TypeString "context.Context"))}}
    {{ .Name.Class }}  {{ .TypeString }}
    {{- end }}
    {{- end }}
}

type {{.Name.Class}}Response struct {
    {{- range .Response }}
    {{- if (and (ne .TypeString "error") (ne .TypeString "context.Context"))}}
    {{ .Name.Class }}  {{ .TypeString }}
    {{- end }}
    {{- end }}
}

{{ .FunCode }}

    {{ end }}
