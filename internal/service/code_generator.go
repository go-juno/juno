package service

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	"github.com/go-juno/juno/internal/constant"
	"github.com/go-juno/juno/internal/util/command"
	"github.com/go-juno/juno/internal/util/generator"
	"github.com/go-juno/juno/internal/util/util"
	"github.com/go-juno/juno/static"
	"golang.org/x/xerrors"
)

type ServiceTplParam struct {
	Mod   string
	Class string
	Camel string
	Snake string // Add Snake field
}

type ServiceWireTplParam struct {
	Mod         string
	ServiceList []*Service
}

// 生成新的service
func GeneratorService(mod, name string) (err error) {
	g, err := generator.NewGenerator(name, constant.ServiceDirPath, mod)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	if g.IsExistsFile() {
		log.Printf("service:%s already exists", name)
		return
	}

	camel, class, _, _ := util.TransformName(name)
	stp := &ServiceTplParam{
		Mod:   mod,
		Class: class,
		Camel: camel,
	}
	tpl, err := template.New("service").Parse(static.ServiceTpl)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	b := bytes.NewBuffer([]byte{})
	err = tpl.Execute(b, stp)
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

// 对更新wire文件
func WireService(mod string) (err error) {
	g, err := generator.NewGenerator("wire", constant.ServiceDirPath, mod)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	// 先解析文件中已存在的包和service
	w, err := ParseServiceWire(g.GetPath())
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	serviceWire := &ServiceWireTplParam{
		Mod:         mod,
		ServiceList: w,
	}

	tpl, err := template.New("s").Parse(static.ServiceWireTpl)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	b := bytes.NewBuffer([]byte{})
	err = tpl.Execute(b, serviceWire)
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

// GenerateHttpRelatedFiles encapsulates the logic for generating HTTP-related files.
func GenerateHttpRelatedFiles() (err error) {
	mod, err := util.GetMod()
	if err != nil {
		return xerrors.Errorf("%w", err)
	}

	// Update service wire
	err = WireService(mod)
	if err != nil {
		return xerrors.Errorf("%w", err)
	}

	// Update endpoint wire
	err = WireEndpoint(mod)
	if err != nil {
		return xerrors.Errorf("%w", err)
	}

	// Create handle
	err = GenerateHandle(mod)
	if err != nil {
		return xerrors.Errorf("%w", err)
	}

	// Generate HTTP content
	err = GenerateHttp(mod)
	if err != nil {
		return xerrors.Errorf("%w", err)
	}

	// Update import information
	err = command.GoimportsCode()
	if err != nil {
		return xerrors.Errorf("%w", err)
	}

	err = command.RunWire()
	if err != nil {
		return xerrors.Errorf("%w", err)
	}
	return nil
}

// GenerateServiceRelatedFiles encapsulates the logic for generating service-related files.
func GenerateServiceRelatedFiles(serviceName, kind string) (err error) {
	mod, err := util.GetMod()
	if err != nil {
		return xerrors.Errorf("%w", err)
	}

	switch kind {
	case "mongo":
		err = GenerateMongoService(mod, serviceName)
		if err != nil {
			return xerrors.Errorf("%w", err)
		}
	default: // Default or empty kind will generate standard service
		if serviceName != "" {
			err = GeneratorService(mod, serviceName)
			if err != nil {
				return xerrors.Errorf("%w", err)
			}
		}
	}

	// These steps are common for all service types
	err = WireService(mod)
	if err != nil {
		return xerrors.Errorf("%w", err)
	}

	err = WireEndpoint(mod)
	if err != nil {
		return xerrors.Errorf("%w", err)
	}

	err = command.RunWire()
	if err != nil {
		return xerrors.Errorf("%w", err)
	}

	err = command.GoimportsCode()
	if err != nil {
		return xerrors.Errorf("%w", err)
	}
	return nil
}

// GenerateMongoService generates a service with MongoDB  operations.
func GenerateMongoService(mod, name string) (err error) {
	camel, class, snake, _ := util.TransformName(name)

	// 1. Generate MongoDB Model
	modelPath := filepath.Join(util.GetPwd(), "internal", "model")
	modelGenerator, err := generator.NewGenerator(name, modelPath, mod)
	if err != nil {
		return xerrors.Errorf("%w", err)
	}
	if modelGenerator.IsExistsFile() {
		log.Printf("model:%s already exists", name)
		// Optionally, decide to overwrite or skip
	}

	modelTplParams := &ServiceTplParam{
		Mod:   mod,
		Class: class,
		Camel: camel,
	}
	modelTpl, err := template.New("mongo_model").Parse(static.MongoModelTpl)
	if err != nil {
		return xerrors.Errorf("%w", err)
	}
	modelBuffer := bytes.NewBuffer([]byte{})
	err = modelTpl.Execute(modelBuffer, modelTplParams)
	if err != nil {
		return xerrors.Errorf("%w", err)
	}
	modelGenerator.SetContent(modelBuffer.String())
	err = modelGenerator.WriteToFile()
	if err != nil {
		return xerrors.Errorf("%w", err)
	}

	// 2. Generate MongoDB  Service
	servicePath := constant.ServiceDirPath
	serviceGenerator, err := generator.NewGenerator(name, servicePath, mod)
	if err != nil {
		return xerrors.Errorf("%w", err)
	}
	if serviceGenerator.IsExistsFile() {
		log.Printf("service:%s already exists", name)
		// Optionally, decide to overwrite or skip
	}

	serviceTplParams := &ServiceTplParam{
		Mod:   mod,
		Class: class,
		Camel: camel,
		Snake: snake, // Snake case for collection name
	}
	serviceTpl, err := template.New("mongo_service").Parse(static.MongoServiceTpl)
	if err != nil {
		return xerrors.Errorf("%w", err)
	}
	serviceBuffer := bytes.NewBuffer([]byte{})
	err = serviceTpl.Execute(serviceBuffer, serviceTplParams)
	if err != nil {
		return xerrors.Errorf("%w", err)
	}
	serviceGenerator.SetContent(serviceBuffer.String())
	err = serviceGenerator.WriteToFile()
	if err != nil {
		return xerrors.Errorf("%w", err)
	}

	return nil
}

type Funcs struct {
	Name   string
	Path   string
	Method string
}

type EndpointFunc struct {
	Mod   string
	Funcs []*Funcs
}

// 解析 endpoint
func parseEndpoint(path string, mod string) (endpointFunc *EndpointFunc, err error) {
	fset := token.NewFileSet()
	dir, err := parser.ParseDir(fset, path, nil, parser.ParseComments)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	endpointFunc = &EndpointFunc{
		Mod:   mod,
		Funcs: []*Funcs{},
	}
	for _, pkg := range dir {
		for _, f := range pkg.Files {
			for _, decl := range f.Decls {
				genDecl, ok := decl.(*ast.FuncDecl)
				if ok {
					fun := &Funcs{
						Name:   genDecl.Name.Name,
						Path:   "",
						Method: "",
					}
					if genDecl.Doc != nil {
						for _, comment := range genDecl.Doc.List {
							if strings.HasPrefix(comment.Text, "// @Router") {
								routerSplit := strings.Split(comment.Text, " ")
								if len(routerSplit) >= 4 {
									fun.Path = routerSplit[2]
									fun.Method = strings.ReplaceAll(strings.ReplaceAll(routerSplit[3], "[", ""), "]", "")
								}
							}

						}
						if fun.Method == "" {
							continue
						}
						endpointFunc.Funcs = append(endpointFunc.Funcs, fun)
					}

				}
			}
		}
	}
	sort.Slice(endpointFunc.Funcs, func(i, j int) bool {
		return endpointFunc.Funcs[i].Name < endpointFunc.Funcs[j].Name
	})

	return
}

func GenerateHttp(mod string) (err error) {
	// 解析所有的endpoint文件
	ef, err := parseEndpoint(constant.EndpointDirPath, mod)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}

	var g *generator.Generator
	g, err = generator.NewGenerator("http", constant.HttpDirPath, mod)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	var tpl *template.Template
	tpl, err = template.New("s").Parse(static.HttpTpl)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	b := bytes.NewBuffer([]byte{})
	err = tpl.Execute(b, ef)
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

func GenerateHandle(mod string) (err error) {
	var g *generator.Generator
	g, err = generator.NewGenerator("handle", constant.HttpDirPath, mod)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	if g.IsExistsFile() {
		return
	}

	var tpl *template.Template
	tpl, err = template.New("s").Parse(static.HandleTpl)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	b := bytes.NewBuffer([]byte{})
	err = tpl.Execute(b, mod)
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
