package generator

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/go-juno/juno/internal/constant"
	"github.com/go-juno/juno/pkg/util"
	"golang.org/x/xerrors"
)

type Generator struct {
	name     string
	path     string
	mod      string
	filePath string
	buf      bytes.Buffer // Accumulated output.
}

func NewGenerator(name, path, mod string) (g *Generator, err error) {
	path = filepath.Join(util.GetPwd(), path)
	_, _, snake, _ := util.TransformName(name)
	// 判断文件是否存在, 如果存在,则需要读取文件内容
	filePath := filepath.Join(path, fmt.Sprintf("%s.go", snake))
	g = &Generator{
		name:     name,
		path:     path,
		mod:      mod,
		filePath: filePath,
		buf:      bytes.Buffer{},
	}

	ok := g.IsExistsFile()
	if ok {
		var content []byte
		content, err = util.ReadAll(filePath)
		if err != nil {
			err = xerrors.Errorf("%w", err)
			return
		}
		g.buf = *bytes.NewBuffer(content)
	}
	return
}

func (g *Generator) IsExistsFile() (ok bool) {
	var err error
	_, err = os.Stat(g.filePath)
	if err != nil {
		if !os.IsExist(err) {
			return
		} else {
			return
		}

	}
	ok = true
	return
}

func (g *Generator) Contains(content string) (ok bool) {
	ok = bytes.Contains(g.buf.Bytes(), []byte(content))
	return
}

func (g *Generator) Printf(format string, args ...interface{}) {
	fmt.Fprintf(&g.buf, format, args...)
}

func (g *Generator) Replace(old string, new string) {
	src := bytes.ReplaceAll(g.buf.Bytes(), []byte(old), []byte(new))
	g.buf = *bytes.NewBuffer(src)
}

func (g *Generator) format() (src []byte, err error) {
	camel, class, snake, hyphen := util.TransformName(g.name)
	src = bytes.ReplaceAll(g.buf.Bytes(), constant.TplHyphen, []byte(hyphen))
	src = bytes.ReplaceAll(src, constant.TplMod, []byte(g.mod))
	src = bytes.ReplaceAll(src, constant.TplCamel, []byte(camel))
	src = bytes.ReplaceAll(src, constant.TplClass, []byte(class))
	src = bytes.ReplaceAll(src, constant.TplSnake, []byte(snake))
	c := src

	src, err = format.Source(src)
	if err != nil {
		err = ioutil.WriteFile(g.filePath+".tpl", c, 0644)
		if err != nil {
			err = xerrors.Errorf("%w", err)
			return
		}
		err = xerrors.Errorf("%w", err)
		return
	}
	return
}

func (g *Generator) WriteToFile() (err error) {
	src, err := g.format()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	//先创建文件夹
	err = util.Mkdir(g.path)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	err = ioutil.WriteFile(g.filePath, src, 0644)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	return
}
