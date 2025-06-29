package cmd

import (
	"fmt"
	"log"

	"github.com/go-juno/juno/internal/endpoint"
	"github.com/spf13/cobra"
)

var serviceCmd = &cobra.Command{
	Use:   "service [service-name]",
	Short: "Create a new service",
	Long:  `Create a new service with a default structure.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		serviceName := args[0]
		kind, _ := cmd.Flags().GetString("kind")
		fmt.Printf("Creating new service: %s (kind: %s)\n", serviceName, kind)

		req := &endpoint.CreateServiceRequest{
			Name: serviceName,
			Kind: kind,
		}
		err := endpoint.CreateServiceEndpoint(req)
		if err != nil {
			log.Printf("err: %+v", err)
		}
	},
}

func init() {
	serviceCmd.Flags().String("kind", "", "Specify the kind of service to create (e.g., mongo-crud)")
	rootCmd.AddCommand(serviceCmd)
}
