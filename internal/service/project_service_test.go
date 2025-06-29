package service

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCreateProject(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir, err := os.MkdirTemp("", "test_project_")
	if err != nil {
		t.Fatalf("failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tmpDir) // Clean up after the test

	// Change to the temporary directory to simulate project creation
	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get current working directory: %v", err)
	}
	if err := os.Chdir(tmpDir); err != nil {
		t.Fatalf("failed to change to temp directory: %v", err)
	}
	defer os.Chdir(originalDir) // Change back after the test

	projectName := "test-app"
	service := NewProjectService()
	err = service.CreateProject(projectName)
	if err != nil {
		t.Fatalf("CreateProject failed: %v", err)
	}

	// Verify that the project directory was created
	projectPath := filepath.Join(tmpDir, projectName)
	if _, err := os.Stat(projectPath); os.IsNotExist(err) {
		t.Errorf("project directory '%s' was not created", projectPath)
	}

	// Verify some key files exist
	expectedFiles := []string{
		filepath.Join(projectPath, "main.go"),
		filepath.Join(projectPath, "go.mod"),
		filepath.Join(projectPath, "README.md"),
		filepath.Join(projectPath, ".gitignore"),
		filepath.Join(projectPath, "Dockerfile"),
		filepath.Join(projectPath, ".air.conf"),
		filepath.Join(projectPath, ".pre-commit-config.yaml"),
	}

	for _, file := range expectedFiles {
		if _, err := os.Stat(file); os.IsNotExist(err) {
			t.Errorf("expected file '%s' was not created", file)
		}
	}

	// Verify some key directories exist
	expectedDirs := []string{
		filepath.Join(projectPath, "api"),
		filepath.Join(projectPath, "cmd"),
		filepath.Join(projectPath, "configs"),
		filepath.Join(projectPath, "internal"),
		filepath.Join(projectPath, "pkg"),
	}

	for _, dir := range expectedDirs {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			t.Errorf("expected directory '%s' was not created", dir)
		}
	}
}
