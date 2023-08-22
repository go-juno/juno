package service

import (
	"bytes"
	"log"
	"text/template"

	"github.com/go-juno/juno/internal/constant"
	"github.com/go-juno/juno/pkg/generator"
	"github.com/go-juno/juno/pkg/util"
	"github.com/go-juno/juno/static"
	"golang.org/x/xerrors"
)

type ServiceTplParam struct {
	Mod   string
	Class string
	Camel string
}

type ServiceWireTplParam struct {
	Mod         string
	ServiceList []*Service
}

// 生成新的service
func GeneratorService(mod, name string) (err error) {
	g, err := generator.NewGenerator(name, constant.ServiceDirPath, mod)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	if g.IsExistsFile() {
		log.Printf("service:%s already exists", name)
		return
	}

	camel, class, _, _ := util.TransformName(name)
	stp := &ServiceTplParam{
		Mod:   mod,
		Class: class,
		Camel: camel,
	}
	tpl, err := template.New("service").Parse(static.ServiceTpl)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	b := bytes.NewBuffer([]byte{})
	err = tpl.Execute(b, stp)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}

	g.SetContent(b.String())
	err = g.WriteToFile()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	return
}

// 对更新wire文件
func WireService(mod, name string) (err error) {
	g, err := generator.NewGenerator("wire", constant.ServiceDirPath, mod)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	// 先解析文件中已存在的包和service
	w, err := ParseServiceWire(g.GetPath())
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	serviceWire := &ServiceWireTplParam{
		Mod:         mod,
		ServiceList: w,
	}

	tpl, err := template.New("s").Parse(static.ServiceWireTpl)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	b := bytes.NewBuffer([]byte{})
	err = tpl.Execute(b, serviceWire)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	g.SetContent(b.String())
	err = g.WriteToFile()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	return
}
