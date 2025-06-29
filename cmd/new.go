package cmd

import (
	"fmt"

	"github.com/go-juno/juno/internal/endpoint"
	"github.com/go-juno/juno/internal/service"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new [project-name]",
	Short: "Create a new project",
	Long:  `Create a new project with a default structure.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		fmt.Printf("Creating new project: %s\n", projectName)

		// Manual dependency injection
		projectService := service.NewProjectService()
		projectEndpoint := endpoint.NewProjectEndpoint(projectService)

		req := &endpoint.CreateProjectRequest{
			Name: projectName,
		}

		if err := projectEndpoint.CreateProject(req); err != nil {
			fmt.Printf("Error creating project: %v\n", err)
		} else {
			fmt.Printf("Project '%s' created successfully.\n", projectName)
		}
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
