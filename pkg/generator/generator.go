package generator

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-juno/juno/internal/constant"
	"github.com/go-juno/juno/pkg/util"
	"golang.org/x/xerrors"
)

type Generator struct {
	name string
	path string
	buf  bytes.Buffer // Accumulated output.
}

func NewGenerator(name string, path string) (g *Generator, err error) {
	g = &Generator{
		name: name,
		path: path,
		buf:  bytes.Buffer{},
	}
	//先创建文件夹
	err = util.Mkdir(path)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	// 判断文件是否存在, 如果存在,则需要读取文件内容
	outputName := filepath.Join(constant.ServiceDirPath, name)
	ok, err := util.IsExistsFile(outputName)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	if ok {
		var content []byte
		content, err = util.ReadAll(outputName)
		if err != nil {
			err = xerrors.Errorf("%w", err)
			return
		}
		_, err = g.buf.Write(content)
		if err != nil {
			err = xerrors.Errorf("%w", err)
			return
		}
	}
	return
}

func (g *Generator) IsExistsFile() (ok bool, err error) {
	camel := util.CamelString(g.name)
	filepath.Join()
	_, err = os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if !os.IsExist(err) {
			err = nil
			return
		} else {
			err = xerrors.Errorf("%w", err)
			return
		}

	}
	ok = true
	return
}

func (g *Generator) Printf(format string, args ...interface{}) {
	fmt.Fprintf(&g.buf, format, args...)
}

func (g *Generator) format() (src []byte, err error) {
	src, err = format.Source(g.buf.Bytes())
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	return
}

func (g *Generator) WriteToFile() (err error) {
	bytes.ReplaceAll(g.buf.Bytes(), []byte(constant.TplHyphen), []byte("\r\n"))
	tpl = strings.ReplaceAll(content, constant.TplHyphen, hyphen)
	tpl = strings.ReplaceAll(tpl, constant.TplMod, mod)
	tpl = strings.ReplaceAll(tpl, constant.TplCamel, camel)
	tpl = strings.ReplaceAll(tpl, constant.TplClass, class)
	tpl = strings.ReplaceAll(tpl, constant.TplSnake, snake)
	var src []byte
	src, err = g.format()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	err = ioutil.WriteFile(g.outputName, src, 0644)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	return
}
