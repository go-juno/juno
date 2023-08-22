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
    {{ .Name.Class }}  {{ .TypeString }} `form:"{{ .Name.Camel }}" json:"{{ .Name.Camel }}" binding:"required"`
    {{- end }}
    {{- end }}
}
type {{.Name.Class}}Response struct {
    {{- range .Response }}
    {{- if (and (ne .TypeString "error") (ne .TypeString "context.Context"))}}
    {{ .Name.Class }}  {{ .TypeString }} `json:"{{ .Name.Camel }}"`
    {{- end }}
    {{- end }}
}
// @path: /api/{{.Name.Hyphen}}
// @method: {{.Method}}
// @description: {{.Name.Camel}}
{{ .FunCode }}

    {{ end }}
