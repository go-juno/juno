package endpoint

import "github.com/go-juno/juno/internal/service"

type ProjectEndpoint struct {
	projectService *service.ProjectService
}

func NewProjectEndpoint(projectService *service.ProjectService) *ProjectEndpoint {
	return &ProjectEndpoint{
		projectService: projectService,
	}
}

type CreateProjectRequest struct {
	Name string
}

func (e *ProjectEndpoint) CreateProject(req *CreateProjectRequest) error {
	return e.projectService.CreateProject(req.Name)
}
