package util

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"golang.org/x/xerrors"
)

var structRegular = `type(.*)Request struct {([\s\w]+)}`
var fieldRegular = `\w+\s+\w+`

func GenField(structName string, content string) (structJson string, err error) {

	structJson = fmt.Sprintf(`type %s struct{
		%{field}
	}	
	`, structName)

	reg, err := regexp.Compile(fieldRegular)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	infoList := reg.FindAllStringSubmatch(content, -1)
	fieldStringList := make([]string, len(infoList))
	for index, info := range infoList {
		fieldString := info[0]
		var name, filedType string
		fieldList := strings.Split(fieldString, " ")
		for _, field := range fieldList {
			d := strings.TrimSpace(field)
			if d != "" {
				if name == "" {
					name = d
				} else {
					filedType = d
				}
			}

		}
		fieldStringList[index] = fmt.Sprintf("  %s %s", name, filedType)
	}

	return
}

func GetRequestStruct(content string) (err error) {
	reg, err := regexp.Compile(structRegular)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	info := reg.FindAllStringSubmatch(content, -1)
	for _, i := range info {
		err = GenField(i[2])
		if err != nil {
			log.Println("err", err)
		}
	}
	return
}
