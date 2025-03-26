package service

import (
	"go/ast"
	"go/parser"
	"go/token"
	"sort"
	"strings"

	"golang.org/x/xerrors"
)

type Service struct {
	Name        *Name
	ServiceName string
	NewFuncName string
}

func ParseServiceWire(path string) (w []*Service, err error) {
	fset := token.NewFileSet()
	dir, err := parser.ParseDir(fset, path, nil, parser.AllErrors)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	w = make([]*Service, 0)
	for _, pkg := range dir {
		for _, f := range pkg.Files {
			service := &Service{}
			for _, decl := range f.Decls {
				funcDecl, ok := decl.(*ast.FuncDecl)
				if ok {
					if funcDecl.Name.Obj != nil {
						if funcDecl.Name.Obj.Kind == ast.Fun {
							if funcDecl.Type.Results == nil {
								continue
							}
							service.NewFuncName = funcDecl.Name.Name
							resultType, ok := funcDecl.Type.Results.List[0].Type.(*ast.Ident)
							if ok {
								lastIndex := strings.LastIndex(resultType.Name, "Service")
								if lastIndex != -1 {
									name := resultType.Name[:lastIndex]
									service.ServiceName = resultType.Name
									service.Name = NewName(name)
									w = append(w, service)
								}

							}
						}
					}

				}
			}
		}

	}
	sort.Slice(w, func(i, j int) bool {
		return w[i].ServiceName < w[j].ServiceName
	})
	return

}
