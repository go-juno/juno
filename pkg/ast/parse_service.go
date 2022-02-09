package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"path/filepath"

	"golang.org/x/tools/go/packages"
)

func parsePackage() {
	cfg := &packages.Config{
		Mode:  packages.NeedSyntax,
		Tests: false,
		Dir:   "/Users/dz0400145/my/kit-service/internal/service",
	}
	pkgs, err := packages.Load(cfg)
	if err != nil {
		log.Fatal(err)
	}
	// log.Println("pkgs[0]", pkgs)
	// fset := token.NewFileSet()
	// ast.Print(fset, pkgs[0].Syntax[1])
	for _, file := range pkgs[0].Syntax {
		log.Println("file", file.Name)
	}
}

// func genDecl(node ast.Node) bool {
// 	decl, ok := node.(*ast.GenDecl)
// 	if !ok {
// 		return true
// 	}
// 	for _, spec := range decl.Specs {
// 		vspec := spec.(*ast.ValueSpec) // Guaranteed to succeed as this is CONST.
// 		if vspec.Type == nil && len(vspec.Values) > 0 {
// 			// "X = 1". With no type but a value. If the constant is untyped,
// 			// skip this vspec and reset the remembered type.
// 			typ = ""

// 			// If this is a simple type conversion, remember the type.
// 			// We don't mind if this is actually a call; a qualified call won't
// 			// be matched (that will be SelectorExpr, not Ident), and only unusual
// 			// situations will result in a function call that appears to be
// 			// a type conversion.
// 			ce, ok := vspec.Values[0].(*ast.CallExpr)
// 			if !ok {
// 				continue
// 			}
// 			id, ok := ce.Fun.(*ast.Ident)
// 			if !ok {
// 				continue
// 			}
// 			typ = id.Name
// 		}
// 		if vspec.Type != nil {
// 			// "X T". We have a type. Remember it.
// 			ident, ok := vspec.Type.(*ast.Ident)
// 			if !ok {
// 				continue
// 			}
// 			typ = ident.Name
// 		}
// 		if typ != f.typeName {
// 			// This is not the type we're looking for.
// 			continue
// 		}
// 		// We now have a list of names (from one line of source code) all being
// 		// declared with the desired type.
// 		// Grab their names and actual values and store them in f.values.
// 		for _, name := range vspec.Names {
// 			if name.Name == "_" {
// 				continue
// 			}
// 			// This dance lets the type checker find the values for us. It's a
// 			// bit tricky: look up the object declared by the name, find its
// 			// types.Const, and extract its value.
// 			obj, ok := f.pkg.defs[name]
// 			if !ok {
// 				log.Fatalf("no value for constant %s", name)
// 			}
// 			info := obj.Type().Underlying().(*types.Basic).Info()
// 			if info&types.IsString == 0 {
// 				log.Fatalf("can't handle non-string constant type %s", typ)
// 			}
// 			value := obj.(*types.Const).Val() // Guaranteed to succeed as this is CONST.
// 			if value.Kind() != constant.String {
// 				log.Fatalf("can't happen: constant is not an string %s", name)
// 			}
// 			v := Value{
// 				originalName: name.Name,
// 				value:        constant.StringVal(value),
// 			}
// 			f.values = append(f.values, v)
// 		}
// 	}
// 	return false
// }

type FieldType string

const (
	FieldTypeCommon FieldType = "common"
	FieldTypeStruct FieldType = "struct"
)

type Field struct {
	Name       string
	TypeString string
	FieldType  FieldType
	Field      []*Field
}

type Func struct {
	Name     string
	Request  []*Field
	Response []*Field
}

type Parser struct {
	Funcs    []*Func
	Packages []string
}

func (f *Field) ParseFeild(expr ast.Expr) {
	switch v := expr.(type) {
	case *ast.StarExpr:
		f.TypeString = fmt.Sprintf("%s*", f.TypeString)
		f.ParseFeild(v.X)
	case *ast.SelectorExpr:
		var prefix string
		ident, ok := v.X.(*ast.Ident)
		if ok {
			prefix = ident.Name
		}
		f.TypeString = fmt.Sprintf("%s%s.%s", f.TypeString, prefix, v.Sel.Name)
		f.FieldType = FieldTypeStruct
	case *ast.Ident:
		if v.Obj != nil {
			f.FieldType = FieldTypeStruct
		} else {
			f.FieldType = FieldTypeCommon
		}
		f.TypeString = fmt.Sprintf("%s%s", f.TypeString, v.Name)
	case *ast.ArrayType:
		f.TypeString = fmt.Sprintf("%s[]", f.TypeString)
		f.ParseFeild(v.Elt)
	}
}

func ParseFile(path, name string) (p *Parser) {

	p = &Parser{}
	fset := token.NewFileSet()
	filename := filepath.Join(path, fmt.Sprintf("%s.go", name))
	f, err := parser.ParseFile(fset, filename, nil, parser.AllErrors)
	if err != nil {
		log.Println(err)
		return
	}
	ast.Print(fset, f)

	for _, decl := range f.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if ok {
			if genDecl.Tok == token.TYPE {
				for _, spec := range genDecl.Specs {
					typespec, ok := spec.(*ast.TypeSpec)
					if ok {
						interfaceType, ok := typespec.Type.(*ast.InterfaceType)
						if ok {

							// 遍历接口
							for _, method := range interfaceType.Methods.List {
								funcs := &Func{
									Name: method.Names[0].Name,
								}
								funcType, ok := method.Type.(*ast.FuncType)
								if ok {
									// 遍历函数参数
									request := make([]*Field, 0)
									for _, param := range funcType.Params.List {
										field := &Field{
											Name:       param.Names[0].Name,
											TypeString: "",
										}
										field.ParseFeild(param.Type)
										request = append(request, field)
									}
									funcs.Request = request
									// 遍历函数返回值
									response := make([]*Field, 0)
									for _, result := range funcType.Results.List {
										field := &Field{
											Name:       result.Names[0].Name,
											TypeString: "",
										}
										field.ParseFeild(result.Type)
										response = append(response, field)
									}
									funcs.Response = response
								}
								p.Funcs = append(p.Funcs, funcs)

							}
						}
					}
				}
			}
			if genDecl.Tok == token.IMPORT {
				for _, spec := range genDecl.Specs {
					importSpec, ok := spec.(*ast.ImportSpec)
					if ok {
						if importSpec.Name != nil {
							p.Packages = append(p.Packages, fmt.Sprintf("%s %s", importSpec.Name.Name, importSpec.Path.Value))

						} else {
							p.Packages = append(p.Packages, importSpec.Path.Value)

						}

					}

				}
			}

		}
	}
	return
}

func main() {
	path := "/Users/dz0400145/my/kit-service/internal/service"
	name := "earning_summary"
	p := ParseFile(path, name)
	log.Println("package:", p.Packages)
	for _, fu := range p.Funcs {
		log.Println(fu.Name)
		for _, req := range fu.Request {
			log.Println("request:", req.Name, req.TypeString, req.FieldType)
		}
		for _, req := range fu.Response {
			log.Println("response:", req.Name, req.TypeString, req.FieldType)
		}
	}
}
