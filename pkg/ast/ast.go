package ast

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"strings"

	"golang.org/x/xerrors"
)

func GetAstFile(filePath string) (f *ast.File, err error) {
	fset := token.NewFileSet()
	f, err = parser.ParseFile(fset, filePath, nil, parser.AllErrors)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	return
}

type RequestStruct struct {
	Name   string
	Struct string
}
type ResponseStruct struct {
	Name   string
	Struct string
}

func GetStruct(f *ast.File) (requestStruct []*RequestStruct, responseStruct []*ResponseStruct) {
	requestStruct = make([]*RequestStruct, 0)
	responseStruct = make([]*ResponseStruct, 0)
	for _, decl := range f.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if ok {
			for _, spec := range genDecl.Specs {
				typeSpec, ok := spec.(*ast.TypeSpec)
				if ok {

					structName := ""
					if strings.Contains(typeSpec.Name.Name, "Request") {
						structName = strings.ReplaceAll(typeSpec.Name.Name, "Request", "")
					}
					if strings.Contains(typeSpec.Name.Name, "Response") {
						structName = strings.ReplaceAll(typeSpec.Name.Name, "Response", "")
					}
					structString := fmt.Sprintf(`type %s struct {`, structName)
					specType, ok := typeSpec.Type.(*ast.StructType)
					if ok {
						for _, field := range specType.Fields.List {
							typeName := ""
							tp, ok := field.Type.(*ast.Ident)
							if ok {
								typeName = tp.Name
							}

							structString += fmt.Sprintf(`
							%s %s `, field.Names[0].Name, typeName)

						}
					}
					structString += `
		}`
					if strings.Contains(typeSpec.Name.Name, "Request") {
						requestStruct = append(requestStruct, &RequestStruct{
							Name:   typeSpec.Name.Name,
							Struct: structString,
						})
					}
					if strings.Contains(typeSpec.Name.Name, "Response") {
						responseStruct = append(responseStruct, &ResponseStruct{
							Name:   typeSpec.Name.Name,
							Struct: structString,
						})
					}

				}
			}
		}
	}
	return
}

type Func struct {
	Endpoint string
}

func GetFunc(f *ast.File) {
	for _, decls := range f.Decls {
		funcDecl, ok := decls.(*ast.FuncDecl)
		if ok {
			log.Println("funcDecl", funcDecl)
		}
	}
}
