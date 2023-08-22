package main

import (
	"go/parser"
	"go/token"
	"log"
	"path/filepath"
)

func main() {
	fset := token.NewFileSet()
	// 这里取绝对路径，方便打印出来的语法树可以转跳到编辑器
	path, _ := filepath.Abs("/Users/KaiJiang/SourceCode/coding/ep-service/api/http/handle/")
	dir, err := parser.ParseDir(fset, path, nil, parser.AllErrors)
	if err != nil {
		log.Println(err)
		return
	}
	for _, pkg := range dir {
		for key := range pkg.Files {
			log.Println("f", key)
		}
	}
	// 打印语法树
	// ast.Print(fset, f)
}
