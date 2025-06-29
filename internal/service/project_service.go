package service

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-juno/juno/static"
)

type ProjectService struct{}

func NewProjectService() *ProjectService {
	return &ProjectService{}
}

func (s *ProjectService) CreateProject(projectName string) error {
	fmt.Printf("Service: Creating project '%s'\n", projectName)

	// 1. Create the project root directory
	if err := os.MkdirAll(projectName, 0755); err != nil {
		return fmt.Errorf("failed to create project directory: %w", err)
	}

	// 2. Create subdirectories
	dirs := []string{"api", "cmd", "configs", "internal", "pkg"}
	for _, dir := range dirs {
		if err := os.MkdirAll(filepath.Join(projectName, dir), 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}
	}

	// 3. Create files from templates
	templateFiles := map[string]string{
		"main.go":                 static.MainTpl,
		"go.mod":                  static.ModTpl,
		"README.md":               static.ReadmeTpl,
		".gitignore":              static.GitignoreTpl,
		"Dockerfile":              static.DockerfileTpl,
		".air.conf":               static.AirTpl,
		".pre-commit-config.yaml": static.PreCommitTpl,
	}

	for path, content := range templateFiles {
		// Note: go.mod is a special case, we need to replace the module name
		if path == "go.mod" {
			content = fmt.Sprintf(content, projectName)
		}

		filePath := filepath.Join(projectName, path)
		if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
			return fmt.Errorf("failed to write file %s: %w", path, err)
		}
	}

	fmt.Printf("Project '%s' structure created successfully.\n", projectName)
	return nil
}
