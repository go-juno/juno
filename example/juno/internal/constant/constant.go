package constant

import (
	"os"

	"juno/init/config"
)

var (
	Config = config.Config
	//RELEASE
	RELEASE = os.Getenv("RELEASE") == "true"
)
