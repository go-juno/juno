package cmd

import (
	"fmt"
	"log"

	"github.com/go-juno/juno/internal/endpoint"
	"github.com/spf13/cobra"
)

var endpointCmd = &cobra.Command{
	Use:   "endpoint [endpoint-name]",
	Short: "Create a new endpoint",
	Long:  `Create a new endpoint with a default structure.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		endpointName := args[0]
		fmt.Printf("Creating new endpoint: %s\n", endpointName)

		req := &endpoint.CreateEndpointRequest{
			Name: endpointName,
		}
		err := endpoint.CreateEndpointEndpoint(req)
		if err != nil {
			log.Printf("err: %+v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(endpointCmd)
}
