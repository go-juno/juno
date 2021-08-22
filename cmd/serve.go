package cmd

import (
	"log"

	"github.com/go-juno/juno/commands"
	"github.com/go-juno/juno/pkg/cli"
)

func Start() {
	log.SetFlags(log.Llongfile | log.LstdFlags)
	cli.SetName("juno").SetVersion("v1.0.22")
	cli.AddCommand(commands.Cmds...).Run()
}
