package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/go-juno/juno/pkg/util"
)

func main() {
	file, err := os.Open("/Users/dz0400145/coding/miao-recruitment/internal/endpoint/apply.go")
	if err != nil {
		log.Println("err", err)
		return
	}
	defer file.Close()
	var content []byte
	content, err = ioutil.ReadAll(file)
	if err != nil {
		log.Println("err", err)
		return
	}

	err = util.GetRequestStruct(string(content))
	if err != nil {
		log.Println("err", err)
		return
	}
}
