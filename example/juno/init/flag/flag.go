package flag

import (
	"flag"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
	"testing"
)

type em struct{}

var (
	Flagconf string
)

func init() {

	flag.StringVar(&Flagconf, "c", "configs/config.yaml", "config path, eg: -c configs/config.yaml")
	testing.Init()
	flag.Parse()
	if Flagconf == "" {
		flag.PrintDefaults()
		log.Fatal("config path is required")
	}

	if !strings.HasPrefix(Flagconf, "/") {
		pkgList := strings.Split(reflect.TypeOf(em{}).PkgPath(), "/")
		pkgName := pkgList[len(pkgList)-3]
		dir, err := os.Getwd()
		if err != nil {
			log.Fatal("config path is required")
		}
		splitList := strings.Split(dir, pkgName)
		if len(splitList) > 1 {
			Flagconf = fmt.Sprintf("%s%s/%s", splitList[0], pkgName, Flagconf)
		}

	}

}
