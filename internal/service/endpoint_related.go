package service

import (
	"fmt"

	"github.com/go-juno/juno/internal/constant"
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
