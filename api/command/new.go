package command

import (
	"errors"
	"fmt"
	"os"

	"github.com/go-juno/juno/internal/endpoint"
	"github.com/go-juno/juno/pkg/cli"
)

type CreateProjectCommand struct {
	endpoint *endpoint.Endpoints
}

func (t *CreateProjectCommand) NewProject(name, selectType string) {
	t.endpoint.
		fmt.Print(" - Env Installation complete\n")
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	dest := fmt.Sprintf("%s/%s", pwd, name)
	if !logic.CopyPath(srcDir, dest) {
		panic(errors.New("Copy dir failed"))
	}
	fmt.Println(" > ok")

	fmt.Print(" - Processing package name")
	if err := logic.ReplaceAll(dest, fmt.Sprintf("github.com/%s", selectType), name); err != nil {
		panic(errors.New("Replace failed"))
	}
	if err := logic.ReplaceMod(dest); err != nil {
		panic(errors.New("Replace go.mod failed"))
	}
	fmt.Println(" > ok")

	fmt.Printf("Project '%s' is generated", name)
}

func (t *CreateProjectCommand) Main() {}

func NewCreateProjectCommand(endpoint *endpoint.Endpoints) *cli.Command {
	return &cli.Command{
		Name:  "new",
		Short: "\tCreate a project\n",
		RunI: &CreateProjectCommand{
			endpoint: endpoint,
		},
	}
}
