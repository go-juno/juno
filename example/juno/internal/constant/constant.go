package constant

import (
	"os"

	"github.com/go-juno/juno/example/juno/init/config"
)

var (
	Config = config.Config
	//RELEASE
	RELEASE = os.Getenv("RELEASE") == "true"
)
