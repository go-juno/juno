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


// @Summary {{.Name.Class}}
// @Schemes
// @Description {{.Name.Class}}
// @Tags {{$.Name.Hyphen}}
// @Accept json
// @Produce json
// @Param request {{if (ne .Method "GET")}} query {{else}} body {{end}} {{.Name.Class }}Request false "{{.Name.Class }}Request"
// @Success 200 {object} {{.Name.Class}}Response
// @Router /api/{{$.Name.Hyphen}} [{{.Method}}]
{{ .FunCode }}

    {{ end }}
