package main

import (
	"log"

	"github.com/go-juno/juno/pkg/parse"
)

func main() {
	w, err := parse.ParseServiceWire("/Users/KaiJiang/SourceCode/kit-service/internal/service")
	log.Println("err", err)
	log.Println("w", w.Import)
	log.Println("w", w.ServiceName)
}
