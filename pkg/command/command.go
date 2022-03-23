package command

import (
	"os/exec"
	"path/filepath"

	"github.com/go-juno/juno/internal/constant"
	"github.com/go-juno/juno/pkg/util"
	"golang.org/x/xerrors"
)

func RunWire() (err error) {
	wirePath := filepath.Join(util.GetPwd(), constant.WireFilePath)
	cmd := exec.Command("wire", wirePath)
	err = cmd.Run()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	return
}

func FmtCode() (err error) {

	cmd := exec.Command("gofmt", "-w", util.GetPwd())
	err = cmd.Run()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	return
}

func GoimportsCode() (err error) {
	cmd := exec.Command("goimports", "-w", util.GetPwd())
	err = cmd.Run()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	return
}

func GenProto(file string) (err error) {
	cmd := exec.Command("protoc", "--go_out=paths=source_relative:.", "--go-grpc_out=paths=source_relative:.", "--go-grpc_opt=require_unimplemented_servers=false", file)
	err = cmd.Run()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}

	return
}

func GenMod(name string) (err error) {
	cmd := exec.Command("go", "mod", "init", name, util.GetPwd())
	err = cmd.Run()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	return
}
