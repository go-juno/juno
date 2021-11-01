package service

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/go-juno/juno/pkg/util"
	"golang.org/x/xerrors"
)

type Filed struct {
	Name string
	Type string
	Tag  string
}

type Request struct {
	Name      string
	IsList    bool
	FiledList []*Filed
}

type Response struct {
	Name       string
	IsList     bool
	FiledList  []*Filed
	Pagination bool
}

type Method struct {
	Name     string
	Request  *Request
	Response *Response
}

var methodRegular = `func (.*) (.*)Endpoint\(ctx context.Context, request (.*)Request\) \(response (.*)Response, err error\)`
var fieldRegular = `\w+\s+\S+`
var structRequestRegular = `type %sRequest struct {[^}]*}`
var structResponseRegular = `type %sResponse struct {[^}]*}`
var structResponseItemRegular = `type %s struct {[^}]*}`

type ParseService interface {
	GenMethodList(content string) (methodList []*Method, err error)
}

type parseService struct {
}

func genRequestField(content string, structName string) (filedList []*Filed, err error) {

	regular := fmt.Sprintf(structRequestRegular, structName)
	reg, err := regexp.Compile(regular)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	info := reg.FindStringSubmatch(content)

	reg, err = regexp.Compile(fieldRegular)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	infoList := reg.FindAllStringSubmatch(info[0], -1)
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
					_, _, snake, _ := util.TransformName(filed.Name)
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

func genResponseItemField(content string, structName string) (filedList []*Filed, err error) {
	regular := fmt.Sprintf(structResponseItemRegular, structName)
	reg, err := regexp.Compile(regular)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	info := reg.FindStringSubmatch(content)
	if len(info) == 0 {
		return
	}

	reg, err = regexp.Compile(fieldRegular)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	infoList := reg.FindAllStringSubmatch(info[0], -1)
	filedLength := len(infoList) - 2
	filedList = make([]*Filed, filedLength)
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
					_, _, snake, _ := util.TransformName(filed.Name)
					filed.Tag = fmt.Sprintf("`json:\"%s\"`", snake)
				}
			}

		}
		filedList[index] = filed
	}

	return

}

func genResponseField(content string, structName string) (filedList []*Filed, pagination bool, err error) {

	regular := fmt.Sprintf(structResponseRegular, structName)
	reg, err := regexp.Compile(regular)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	info := reg.FindStringSubmatch(content)

	reg, err = regexp.Compile(fieldRegular)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	infoList := reg.FindAllStringSubmatch(info[0], -1)
	filedLength := len(infoList) - 2
	filedList = make([]*Filed, filedLength)
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
					_, _, snake, _ := util.TransformName(filed.Name)
					filed.Tag = fmt.Sprintf("`json:\"%s\"`", snake)

				}
			}

		}
		filedList[index] = filed
	}
	if filedLength == 2 {
		structName := ""
		if filedList[0].Name == "Total" && strings.Contains(filedList[1].Type, "[]*") {
			structName = strings.ReplaceAll(filedList[1].Type, "[]*", "")
			pagination = true

		} else if filedList[1].Name == "Total" && strings.Contains(filedList[0].Type, "[]*") {
			structName = strings.ReplaceAll(filedList[0].Type, "[]*", "")
			pagination = true
		}
		if pagination {
			filedList, err = genResponseItemField(content, structName)
			if err != nil {
				err = xerrors.Errorf("%w", err)
				return
			}
		}
	}

	return
}

func (s *parseService) GenMethodList(content string) (methodList []*Method, err error) {
	reg, err := regexp.Compile(methodRegular)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	info := reg.FindAllStringSubmatch(content, -1)

	methodList = make([]*Method, len(info))
	for index, matchList := range info {

		if len(matchList) != 5 {
			continue
		}

		requestStringsList := strings.Split(matchList[3], "*")
		responseStringsList := strings.Split(matchList[4], "*")

		var requestField, reponseField []*Filed
		var pagination bool

		requestField, err = genRequestField(content, requestStringsList[1])
		if err != nil {
			err = xerrors.Errorf("%w", err)
			return
		}
		reponseField, pagination, err = genResponseField(content, requestStringsList[1])
		if err != nil {
			err = xerrors.Errorf("%w", err)
			return
		}

		request := &Request{
			Name:      requestStringsList[1],
			IsList:    requestStringsList[0] != "",
			FiledList: requestField,
		}

		response := &Response{
			Name:       responseStringsList[1],
			IsList:     responseStringsList[0] != "",
			FiledList:  reponseField,
			Pagination: pagination,
		}

		methodList[index] = &Method{
			Name:     matchList[2],
			Request:  request,
			Response: response,
		}

	}
	return
}

func NewParseService() ParseService {
	return &parseService{}
}
