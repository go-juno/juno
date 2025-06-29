package cmd

import (
	"fmt"
	"log"

	"github.com/go-juno/juno/internal/endpoint"
	"github.com/spf13/cobra"
)

var grpcCmd = &cobra.Command{
	Use:   "grpc [grpc-name]",
	Short: "Create a new grpc",
	Long:  `Create a new grpc with a default structure.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		grpcName := args[0]
		fmt.Printf("Creating new grpc: %s\n", grpcName)

		req := &endpoint.CreateGrpcRequest{
			Name: grpcName,
		}
		err := endpoint.CreateGrpcEndpoint(req)
		if err != nil {
			log.Printf("err: %+v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(grpcCmd)
}
