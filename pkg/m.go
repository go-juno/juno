package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/go-juno/juno/pkg/util"
)

func main() {
	file, err := os.Open("/Users/dz0400145/coding/miao-recruitment/internal/endpoint/interview.go")
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
	var ss string
	ss, err = util.GenAllRequestStruct(string(content))
	if err != nil {
		log.Println("err", err)
		return
	}
	log.Println("ss", ss)
}
