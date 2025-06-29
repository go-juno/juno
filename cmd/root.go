package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "juno",
	Short: "juno is a lightweight Go microservice framework",
	Long:  `juno is a lightweight Go microservice framework, containing a large number of microservice-related frameworks and tools.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
