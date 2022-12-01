package service

import (
	"bytes"
	"path/filepath"
	"text/template"

	"github.com/go-juno/juno/internal/constant"
	"github.com/go-juno/juno/pkg/generator"
	"github.com/go-juno/juno/pkg/util"
	"github.com/go-juno/juno/static"
	"golang.org/x/xerrors"
)

func WireEndpoint(mod, name string) (err error) {
	// g, err := generator.NewGenerator("endpoint", constant.EndpointDirPath, mod)
	// if err != nil {
	// 	err = xerrors.Errorf("%w", err)
	// 	return
	// }
	// camel, class, _, _ := util.TransformName(name)
	// if !g.IsExistsFile() {
	// 	g.Printf("package endpoint\n")
	// 	g.Printf("import (\n")
	// 	g.Printf("\t\"%s/internal/service\"\n", mod)
	// 	g.Printf("\t\"github.com/google/wire\"\n")
	// 	g.Printf("type Endpoints struct {\n")
	// 	g.Printf("}\n")
	// 	g.Printf("func NewEndpoints(\n")
	// 	g.Printf(") *Endpoints {\n")
	// 	g.Printf("\treturn &Endpoints{\n")
	// 	g.Printf("\t}\n")
	// 	g.Printf("}\n")
	// 	g.Printf("var ProviderSet = wire.NewSet(NewEndpoints)\n")
	// } else {
	// 	exitsService := fmt.Sprintf(`%sService`, class)
	// 	g.Contains(exitsService)
	// }
	// structString := fmt.Sprintf("type Endpoints struct {\n\t%s service.%sService", camel, class)
	// paramString := fmt.Sprintf("func NewEndpoints(\n\t%s service.%sService,", camel, class)
	// classString := fmt.Sprintf("return &Endpoints{\n\t\t%s: %s,", camel, camel)
	// g.Replace("type Endpoints struct {", structString)
	// g.Replace("func NewEndpoints(", paramString)
	// g.Replace("return &Endpoints{", classString)
	// err = g.WriteToFile()
	// if err != nil {
	// 	err = xerrors.Errorf("%w", err)
	// 	return
	// }
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
	p, err := ParseFile(path, name)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	tpl, err := template.New("s").Parse(static.EndpointTpl)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	b := bytes.NewBuffer([]byte{})
	err = tpl.Execute(b, p)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	g.SetContent(b.String())
	err = g.WriteToFile()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	return
}
