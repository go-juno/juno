package util

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"golang.org/x/xerrors"
)

var structRequestRegular = `type %sRequest struct {[^}]*}`
var structRequestNameRegular = `request (.*)Request`
var fieldRegular = `\w+\s+\S+`

type Filed struct {
	Name string
	Type string
	Tag  string
}

func GenRequestTransform(structName string, filedList []*Filed, isList bool) (transformString string) {

	transformFieldString := ""
	for _, field := range filedList {
		transformFieldString += fmt.Sprintf("\t\t%s: s.%s,\n", field.Name, field.Name)
	}
	transformString = fmt.Sprintf(`func (s *%s) Transform() *endpoint.%sRequest {
	req := &endpoint.%sRequest{
%s    }
	return req
}
	`, structName, structName, structName, transformFieldString)

	return
}

func GenRequestField(content string, structName string) (filedList []*Filed, err error) {
	reg, err := regexp.Compile(fieldRegular)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	infoList := reg.FindAllStringSubmatch(content, -1)
	filedList = make([]*Filed, len(infoList)-2)
	for index, info := range infoList[2:] {
		filed := &Filed{}
		fieldStringList := strings.Split(info[0], " ")

		for _, fieldString := range fieldStringList {
			d := strings.TrimSpace(fieldString)
			if d != "" {
				if filed.Name == "" {
					filed.Name = d
				} else {
					filed.Type = d
					_, _, snake, _ := TransformName(filed.Name)
					var binding string
					if strings.Contains(d, "int") {
						binding = `binding:"required,min=1"`
					} else if strings.Contains(d, "string") {
						binding = `binding:"required,max=100"`
					} else if strings.Contains(d, "LocalTime") {
						binding = `binding:"required" time_format:"2006-01-02 15:04:05"`
					} else if strings.Contains(d, "LocalDate") {
						binding = `binding:"required" time_format:"2006-01-02"`
					} else {
						binding = `binding:"required"`
					}
					filed.Tag = fmt.Sprintf("`form:\"%s\" json:\"%s\" %s`", snake, snake, binding)

				}
			}

		}
		filedList[index] = filed
	}
	return
}

func GenRequestStruct(content, structName string, isList bool) (structString string, err error) {

	regular := fmt.Sprintf(structRequestRegular, structName)
	reg, err := regexp.Compile(regular)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	info := reg.FindStringSubmatch(content)
	var fieldList []*Filed

	fieldList, err = GenRequestField(info[0], structName)
	if err != nil {
		log.Println("err", err)
	}
	fieldStringList := make([]string, len(fieldList))

	for index, field := range fieldList {
		fieldStringList[index] = fmt.Sprintf("\t%s %s %s", field.Name, field.Type, field.Tag)
	}

	transformString := GenRequestTransform(structName, fieldList, isList)

	structString = fmt.Sprintf(`type %s struct {
%s
}

%s
`, structName, strings.Join(fieldStringList, "\n"), transformString)

	return
}

func GenAllRequestStruct(content string) (structString string, err error) {
	reg, err := regexp.Compile(structRequestNameRegular)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	info := reg.FindAllStringSubmatch(content, -1)
	for _, i := range info {
		stringsList := strings.Split(i[1], "*")
		structName := stringsList[1]
		isList := stringsList[0] == ""
		var ss string
		ss, err = GenRequestStruct(content, structName, isList)
		if err != nil {
			err = xerrors.Errorf("%w", err)
			return
		}
		structString += ss
	}
	return
}
