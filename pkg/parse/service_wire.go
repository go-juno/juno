package parse

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
	"sort"

	"golang.org/x/xerrors"
)

type ServiceWire struct {
	ServiceName []string
	Import      []string
}

func ParseServiceWire(path string) (w *ServiceWire, err error) {
	fset := token.NewFileSet()
	filename := filepath.Join(path, "service.go")
	f, err := parser.ParseFile(fset, filename, nil, parser.AllErrors)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	w = &ServiceWire{
		ServiceName: []string{},
		Import:      []string{},
	}

	for _, decl := range f.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if ok {
			// 找到所有导入的包

			if genDecl.Tok == token.IMPORT {
				for _, ipt := range genDecl.Specs {
					importSpec, ok := ipt.(*ast.ImportSpec)
					if ok {
						if importSpec.Name != nil {
							w.Import = append(w.Import, fmt.Sprintf("%s %s", importSpec.Name.Name, importSpec.Path.Value))
						} else {
							w.Import = append(w.Import, importSpec.Path.Value)
						}

					}
				}
			}

			// 找到所有的service
			if genDecl.Tok == token.VAR {
				for _, spec := range genDecl.Specs {
					valueSpec, ok := spec.(*ast.ValueSpec)
					if ok {
						for _, value := range valueSpec.Values {
							callExpr, ok := value.(*ast.CallExpr)
							if ok {

								for _, arg := range callExpr.Args {
									switch v := arg.(type) {
									case *ast.SelectorExpr:
										x, ok := v.X.(*ast.Ident)
										if ok {
											w.ServiceName = append(w.ServiceName, fmt.Sprintf("%s.%s", x.Name, v.Sel.Name))
										}

									case *ast.Ident:
										w.ServiceName = append(w.ServiceName, v.Name)
									}

								}

							}
						}
					}
				}
			}
		}
	}
	sort.Strings(w.Import)
	sort.Strings(w.ServiceName)
	return

}
