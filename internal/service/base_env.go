package service

import (
	"os/exec"

	"golang.org/x/xerrors"
)

type BaseEnvService interface {
	InstallEnv() (err error)
	GetGoEnvPath() (path string, err error)
}

type baseEnvService struct {
}

func (s *baseEnvService) GetGoEnvPath() (path string, err error) {
	envCmd := "go env GOPATH"
	cmd := exec.Command("go", "env", "GOPATH")
	out, err := cmd.CombinedOutput()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	path = string(out[:len(out)-1])
	if path == "" {
		err = xerrors.Errorf("$GOPATH is not configured, see '%s'\n", envCmd)
		return
	}
	return
}

func (s *baseEnvService) InstallEnv() (err error) {
	// 下载protobuf
	exec.Command("go", "get", "github.com/golang/protobuf/proto")
	exec.Command("go", "get", "github.com/golang/protobuf/protoc-gen-go")
	// 下载wire
	exec.Command("go", "get", "wire")
	return
}
func NewBaseEnvService() BaseEnvService {
	return &baseEnvService{}
}
