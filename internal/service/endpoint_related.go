package service

import (
	"fmt"
	"path/filepath"

	"github.com/go-juno/juno/internal/constant"
	"github.com/go-juno/juno/pkg/ast"
	"github.com/go-juno/juno/pkg/generator"
	"github.com/go-juno/juno/pkg/util"
	"golang.org/x/xerrors"
)

func WireEndpoint(mod, name string) (err error) {
	g, err := generator.NewGenerator("endpoint", constant.EndpointDirPath, mod)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	camel, class, _, _ := util.TransformName(name)
	if !g.IsExistsFile() {
		g.Printf("package endpoint\n")
		g.Printf("import (\n")
		g.Printf("\t\"%s/internal/service\"\n", mod)
		g.Printf("\t\"github.com/google/wire\"\n")
		g.Printf("type Endpoints struct {\n")
		g.Printf("}\n")
		g.Printf("func NewEndpoints(\n")
		g.Printf(") *Endpoints {\n")
		g.Printf("\treturn &Endpoints{\n")
		g.Printf("\t}\n")
		g.Printf("}\n")
		g.Printf("var ProviderSet = wire.NewSet(NewEndpoints)\n")
	} else {
		exitsService := fmt.Sprintf(`%sService`, class)
		g.Contains(exitsService)
	}
	structString := fmt.Sprintf("type Endpoints struct {\n\t%s service.%sService", camel, class)
	paramString := fmt.Sprintf("func NewEndpoints(\n\t%s service.%sService,", camel, class)
	classString := fmt.Sprintf("return &Endpoints{\n\t\t%s: %s,", camel, camel)
	g.Replace("type Endpoints struct {", structString)
	g.Replace("func NewEndpoints(", paramString)
	g.Replace("return &Endpoints{", classString)
	err = g.WriteToFile()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	return
}

func GeneratorEndpoint(mod, name string) (err error) {
	g, err := generator.NewGenerator(name, constant.EndpointDirPath, mod)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	if g.IsExistsFile() {
		err = xerrors.Errorf("endpoint:%s already exists", name)
		return
	}
	path := filepath.Join(util.GetPwd(), constant.ServiceDirPath)
	p, err := ast.ParseFile(path, name)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	g.Printf("package endpoint\n")

	//判断是否有package需要inport
	if len(p.Packages) != 0 {
		g.Printf("import (\n")
		for _, packageName := range p.Packages {
			g.Printf("%s\n", packageName)
		}
		g.Printf(")\n")
	}

	err = g.WriteToFile()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	return
}
