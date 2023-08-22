package service

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"

	"github.com/go-juno/juno/pkg/util"
	"golang.org/x/tools/go/packages"
	"golang.org/x/xerrors"
)

type Name struct {
	Class  string
	Camel  string
	Snake  string
	Hyphen string
}

type Field struct {
	Name       Name
	TypeString string
}

type Func struct {
	Name     Name
	Request  []*Field
	Response []*Field
	FunCode  string
	Method   string
}

type Parser struct {
	Name     Name
	Funcs    []*Func
	Packages []string
}

type Pkg struct {
	Files []*ast.File
}

func NewName(name string) (n *Name) {
	camel, class, snake, hyphen := util.TransformName(name)
	n = &Name{
		Class:  class,
		Camel:  camel,
		Snake:  snake,
		Hyphen: hyphen,
	}
	return
}

func NewPkg(dirs []string) (pkg *Pkg, err error) {
	pkg = &Pkg{
		Files: []*ast.File{},
	}
	for _, dir := range dirs {
		cfg := &packages.Config{
			Mode:  packages.NeedSyntax,
			Tests: false,
			Dir:   dir,
		}
		var pkgs []*packages.Package
		pkgs, err = packages.Load(cfg)
		if err != nil {
			err = xerrors.Errorf("%w", err)
			return
		}
		pkg.Files = append(pkg.Files, pkgs[0].Syntax...)
	}
	return
}

func (f *Func) GenerateFunc(name string) (code string) {
	camel, _, _, _ := util.TransformName(name)
	// 写入func
	code = fmt.Sprintf("func (e *Endpoints) %sEndpoint(ctx context.Context, request *%sRequest) (response *%sResponse, err error) {\n", f.Name.Class, f.Name.Class, f.Name.Class)
	//获取请求值变量
	requestParam := ""
	for _, field := range f.Request {
		if field.TypeString == "error" || field.TypeString == "context.Context" {
			continue
		}
		if requestParam != "" {
			requestParam += ", "
		}
		requestParam += "request." + field.Name.Class
	}

	//获取返回值变量
	responseParam := ""
	for _, field := range f.Response {
		if responseParam != "" {
			responseParam += ", "
		}
		responseParam += field.Name.Camel
	}
	if len(f.Response) == 1 {
		code += fmt.Sprintf("\t%s = e.%s.%s(ctx, %s) \n", responseParam, camel, f.Name.Class, requestParam)

	} else {
		code += fmt.Sprintf("\t%s := e.%s.%s(ctx, %s) \n", responseParam, camel, f.Name.Class, requestParam)

	}

	code += "\tif err != nil {\n"
	code += "\t\terr = xerrors.Errorf(\"%w\", err)\n"
	code += "\t\treturn\n"
	code += "\t}\n"

	//返回endpoint
	code += fmt.Sprintf("\tresponse = &%sResponse{\n", f.Name.Class)
	for _, field := range f.Response {
		if field.TypeString == "error" || field.TypeString == "context.Context" {
			continue
		}
		code += fmt.Sprintf("\t\t%s: %s,\n", field.Name.Class, field.Name.Camel)
	}
	code += "\t}\n"

	code += "\treturn \n"
	code += "}\n"
	return

}

func (f *Func) GenerateEmptyFunc(name string) (code string) {
	// 写入func
	code = `// path: /hello
// method: get
// description: hello`
	code += fmt.Sprintf("\nfunc (e *Endpoints) %sEndpoint(ctx context.Context, request *%sRequest) (response *%sResponse, err error) {\n", f.Name.Class, f.Name.Class, f.Name.Class)
	code += "\t// TODO\n"
	code += "\treturn \n"
	code += "}\n"
	return

}

func (f *Field) ParseFeildType(expr ast.Expr) {
	switch v := expr.(type) {
	case *ast.StarExpr:
		f.TypeString = fmt.Sprintf("%s*", f.TypeString)
		f.ParseFeildType(v.X)
	case *ast.SelectorExpr:
		var prefix string
		ident, ok := v.X.(*ast.Ident)
		if ok {
			prefix = ident.Name
		}
		f.TypeString = fmt.Sprintf("%s%s.%s", f.TypeString, prefix, v.Sel.Name)
	case *ast.Ident:
		f.TypeString = fmt.Sprintf("%s%s", f.TypeString, v.Name)
	case *ast.ArrayType:
		f.TypeString = fmt.Sprintf("%s[]", f.TypeString)
		f.ParseFeildType(v.Elt)
	}
}

func (pkg *Pkg) ParseFeild(af *ast.Field) (f *Field) {
	if len(af.Names) == 0 {
		return
	}
	name := NewName(af.Names[0].Name)
	f = &Field{
		Name:       *name,
		TypeString: "",
	}
	f.ParseFeildType(af.Type)
	return f
}

func ParseFile(path, name string) (p *Parser, err error) {
	p = &Parser{}
	fset := token.NewFileSet()
	filename := filepath.Join(path, fmt.Sprintf("%s.go", name))
	ok, _ := util.IsExistsFile(filename)
	if !ok {
		funcs := &Func{
			Name: *NewName("hello"),
		}
		p = &Parser{
			Name: *NewName(name),
			Funcs: []*Func{{
				Name:     *NewName("hello"),
				Request:  []*Field{},
				Response: []*Field{},
				FunCode:  funcs.GenerateEmptyFunc(name),
			}},
			Packages: []string{},
		}
		return
	}
	f, err := parser.ParseFile(fset, filename, nil, parser.AllErrors)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	for _, decl := range f.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if ok {
			if genDecl.Tok == token.TYPE {
				for _, spec := range genDecl.Specs {
					typespec, ok := spec.(*ast.TypeSpec)
					if ok {
						interfaceType, ok := typespec.Type.(*ast.InterfaceType)
						if ok {

							var pkg *Pkg
							pkg, err = NewPkg([]string{path})
							if err != nil {
								err = xerrors.Errorf("%w", err)
								return
							}
							// 遍历接口
							for _, method := range interfaceType.Methods.List {
								funcs := &Func{
									Name: *NewName(method.Names[0].Name),
								}
								funcType, ok := method.Type.(*ast.FuncType)
								if ok {
									// 遍历函数参数
									request := make([]*Field, 0)
									for _, param := range funcType.Params.List {
										feild := pkg.ParseFeild(param)
										if feild != nil {
											request = append(request, feild)
										}

									}
									funcs.Request = request
									// 遍历函数返回值
									response := make([]*Field, 0)
									for _, result := range funcType.Results.List {
										feild := pkg.ParseFeild(result)
										if feild != nil {
											response = append(response, feild)
										}
									}
									funcs.Response = response
									funcs.Method = util.GetMethod(funcs.Name.Camel)
								}
								funcs.FunCode = funcs.GenerateFunc(name)
								p.Funcs = append(p.Funcs, funcs)
								p.GenPackage(pkg)

							}
						}
					}
				}
			}

		}
	}

	return
}

func (p *Parser) GenPackage(pkg *Pkg) {
	for _, f := range pkg.Files {
		for _, decl := range f.Decls {
			genDecl, ok := decl.(*ast.GenDecl)
			if ok {
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
	}

}

func (f *Field) Log(prefix string) {
	if f == nil {
		return
	}
	fmt.Printf("%s%s %s \n", prefix, f.Name, f.TypeString)

}
