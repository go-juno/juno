package service

import (
	"bytes"
	"path/filepath"
	"text/template"

	"github.com/go-juno/juno/internal/constant"
	"github.com/go-juno/juno/pkg/generator"
	"github.com/go-juno/juno/pkg/util"
	"github.com/go-juno/juno/static"
	"golang.org/x/xerrors"
)

func GeneratorEndpoint(mod, name string) (err error) {
	g, err := generator.NewGenerator(name, constant.EndpointDirPath, mod)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	if g.IsExistsFile() {
		err = xerrors.Errorf("endpoint:%s already exists", name)
		return
	}
	_, _, snake, _ := util.TransformName(name)
	path := filepath.Join(util.GetPwd(), constant.ServiceDirPath)
	p, err := ParseFile(path, snake)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	tpl, err := template.New("s").Parse(static.EndpointTpl)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	b := bytes.NewBuffer([]byte{})
	err = tpl.Execute(b, p)
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

func WireEndpoint(mod, name string) (err error) {
	g, err := generator.NewGenerator("wire", constant.EndpointDirPath, mod)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	// 先解析文件中已存在的包和service
	w, err := ParseServiceWire(constant.ServiceDirPath)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	serviceWire := &ServiceWireTplParam{
		Mod:         mod,
		ServiceList: w,
	}

	tpl, err := template.New("s").Parse(static.EndpointWire)
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
