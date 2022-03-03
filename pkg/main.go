package main

import (
	"log"

	"github.com/go-juno/juno/pkg/ast"
)

func main() {
	path := "/Users/admin/my/kit-service/internal/service"
	name := "earning_summary"
	p, err := ast.ParseFile(path, name)
	if err != nil {
		log.Printf("err:%+v", err)
	}
	for _, fu := range p.Funcs {
		for _, req := range fu.Request {
			req.Log()
		}
		for _, res := range fu.Response {
			res.Log()
		}
	}
}
