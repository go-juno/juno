package service

import (
{{ range .Import }}
 {{ . }}
{{- end }}
)



var ProviderSet = wire.NewSet(
{{ range .ServiceName }}
 {{ . }},
{{- end }}
)
