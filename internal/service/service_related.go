package service

import (
	"bytes"
	"fmt"
	"html/template"
	"sort"

	"github.com/go-juno/juno/internal/constant"
	"github.com/go-juno/juno/pkg/generator"
	"github.com/go-juno/juno/pkg/parse"
	"github.com/go-juno/juno/pkg/util"
	"github.com/go-juno/juno/static"
	"golang.org/x/xerrors"
)

type ServiceTplParam struct {
	mod   string
	class string
	camel string
}

// 生成新的service
func GeneratorService(mod, name string) (err error) {
	g, err := generator.NewGenerator(name, constant.ServiceDirPath, mod)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	if g.IsExistsFile() {
		err = xerrors.Errorf("service:%s already exists", name)
		return
	}

	camel, class, _, _ := util.TransformName(name)
	stp := &ServiceTplParam{
		mod:   mod,
		class: class,
		camel: camel,
	}
	tpl, err := template.ParseFS(static.ServiceTpl)
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
	g, err := generator.NewGenerator("service", constant.ServiceDirPath, mod)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}

	var w *parse.ServiceWire
	// 当前service的名称
	sn := fmt.Sprintf("New%sService", util.TitleString(name))

	if g.IsExistsFile() {

		// 先解析文件中已存在的包和service
		w, err = parse.ParseServiceWire(g.GetPath())
		if err != nil {
			err = xerrors.Errorf("%w", err)
			return
		}
		// 加入当前的service
		w.ServiceName = append(w.ServiceName, sn)
		sort.Strings(w.ServiceName)
	} else {
		w = &parse.ServiceWire{
			ServiceName: []string{sn},
			Import:      []string{"github.com/google/wire"},
		}
	}

	tpl, err := template.ParseFS(static.ServiceWire)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	b := bytes.NewBuffer([]byte{})
	err = tpl.Execute(b, w)
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
