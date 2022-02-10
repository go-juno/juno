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
	path := "/Users/dz0400145/my/kit-service/internal/model"
	name := "user"
	fset := token.NewFileSet()
	filename := filepath.Join(path, fmt.Sprintf("%s.go", name))
	f, err := parser.ParseFile(fset, filename, nil, parser.AllErrors)
	if err != nil {
		log.Println(err)
		return
	}
	ast.Print(fset, f)
}
