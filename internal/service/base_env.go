package service

import (
	"os/exec"
	"runtime"
	"strings"

	"golang.org/x/xerrors"
)

type BaseEnvService interface {
	InstallEnv() (err error)
}

type baseEnvService struct {
}

func (s *baseEnvService) InstallEnv() (err error) {
	envCmd := "go env GOPATH"
	cmd := exec.Command("go", "env", "GOPATH")
	out, err := cmd.CombinedOutput()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	goPath := string(out[:len(out)-1])
	if goPath == "" {
		err = xerrors.Errorf("$GOPATH is not configured, see '%s'\n", envCmd)
		return
	}

	dr := ":"
	if runtime.GOOS == "windows" {
		dr = ";"
	}
	if strings.Contains(goPath, dr) {
		err = xerrors.Errorf("$GOPATH cannot have multiple directories, see '%s'\n", envCmd)
		return
	}
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
