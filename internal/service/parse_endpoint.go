package service

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/token"
	"sort"
	"strings"
	"text/template"

	"github.com/go-juno/juno/internal/constant"
	"github.com/go-juno/juno/pkg/generator"
	"github.com/go-juno/juno/static"
	"golang.org/x/xerrors"
)

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
		return endpointFunc.Funcs[i].Path < endpointFunc.Funcs[j].Path
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
