package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"path/filepath"
)

func main() {
	fset := token.NewFileSet()
	filename := filepath.Join("/Users/KaiJiang/SourceCode/kit-service/internal/service", "service.go")
	f, err := parser.ParseFile(fset, filename, nil, parser.AllErrors)
	if err != nil {
		log.Println("err", err)
	}

	ast.Print(fset, f)

	for _, decl := range f.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if ok {
			// 找到所有导入的包

			if genDecl.Tok == token.IMPORT {
				for _, ipt := range genDecl.Specs {
					importSpec, ok := ipt.(*ast.ImportSpec)
					if ok {
						log.Println("111", fmt.Sprintf("%s %s", importSpec.Name.Name, importSpec.Path.Value))
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
											log.Println(222, fmt.Sprintf("%s.%s", x.Name, v.Sel.Name))
										}

									}
									ident, ok := arg.(*ast.Ident)
									if ok {
										log.Println(222, ident.Name)
									}

								}

							}
						}
					}
				}
			}
		}
	}

}
