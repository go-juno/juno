package service

import (
	"fmt"

	"github.com/go-juno/juno/internal/constant"
	"github.com/go-juno/juno/pkg/generator"
	"github.com/go-juno/juno/pkg/util"
	"github.com/go-juno/juno/static"
	"golang.org/x/xerrors"
)

func GeneratorService(mod, name string) (err error) {
	g, err := generator.NewGenerator(name, constant.ServiceDirPath, mod)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	if g.IsExistsFile() {
		err = xerrors.Errorf("service:%s already exists", name)
		return
	}
	g.Printf(static.ServiceTpl)
	err = g.WriteToFile()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	return
}

func WireService(mod, name string) (err error) {
	g, err := generator.NewGenerator("service", constant.ServiceDirPath, mod)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	if !g.IsExistsFile() {
		g.Printf("package service\n")
		g.Printf("import \"github.com/google/wire\"\n")
		g.Printf("var ProviderSet = wire.NewSet(")
		g.Printf(")")
	}

	_, class, _, _ := util.TransformName(name)
	newServiceString := fmt.Sprintf("New%sService,", class)
	g.Replace(")", newServiceString)
	g.Printf(")\n")
	err = g.WriteToFile()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	return
}
