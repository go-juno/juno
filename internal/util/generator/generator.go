package generator

import (
	"fmt"
	"go/format"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-juno/juno/internal/util/util"
	"golang.org/x/xerrors"
)

type Generator struct {
	name     string
	path     string
	mod      string
	filePath string
	content  string
}

func NewGenerator(name, path, mod string) (g *Generator, err error) {
	path = filepath.Join(util.GetPwd(), path)
	_, _, snake, _ := util.TransformName(name)
	filePath := filepath.Join(path, fmt.Sprintf("%s.go", snake))
	g = &Generator{
		name:     name,
		path:     path,
		mod:      mod,
		filePath: filePath,
		content:  "",
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
	ok = strings.Contains(g.content, content)
	return
}

func (g *Generator) format() (src []byte, err error) {
	src, err = format.Source([]byte(g.content))
	if err != nil {
		log.Printf("format err:%+v", err)
		src = []byte(g.content)
		err = nil
		return
	}
	return
}

func (g *Generator) SetContent(content string) {
	g.content = content
}

func (g *Generator) GetPath() string {
	return g.path
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
	err = os.WriteFile(g.filePath, src, 0644)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	return
}
