package cmd

import (
	"log"

	"github.com/go-juno/juno/internal/constant"
	"github.com/go-juno/juno/pkg/cli"
)

func Start() {
	// log.SetFlags(log.Llongfile | log.LstdFlags)
	log.SetFlags(0)
	commands, err := InitServer()
	if err != nil {
		panic(err)
	}

	cli.SetName("juno").SetVersion(constant.Version)
	cli.AddCommand(commands...).Run()
}
