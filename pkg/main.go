package main

import (
	"log"

	"github.com/go-juno/juno/pkg/ast"
)

func main() {
	filePath := "/Users/joker/coding/miao-promotion/internal/endpoint/appraise_committee.go"
	f, err := ast.GetAstFile(filePath)
	if err != nil {
		panic(err)
	}
	reqList, resList := ast.GetStruct(f)
	for _, s := range reqList {
		log.Println("req", s)
	}
	for _, s := range resList {
		log.Println("res", s)
	}
}
