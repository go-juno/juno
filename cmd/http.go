package cmd

import (
	"fmt"
	"log"

	"github.com/go-juno/juno/internal/endpoint"
	"github.com/spf13/cobra"
)

var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Create a new http",
	Long:  `Create a new http with a default structure.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Creating new http")

		req := &endpoint.CreateHttpRequest{}
		err := endpoint.CreateHttpEndpoint(req)
		if err != nil {
			log.Printf("err: %+v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(httpCmd)
}
