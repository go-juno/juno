package static

import (
	_ "embed"
)

//go:embed api-doc/index.html
var ApiDoc string
