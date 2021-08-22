package service

import (
	"io"
	"os"
	"strings"

	"github.com/go-juno/juno/pkg/util"
	"golang.org/x/xerrors"
	"gorm.io/gorm"
)

type FileService interface {
	IsExitsDir(dir string) (ok bool, err error)
	IsExistsFile(path string) (ok bool, err error)
	Mkdir(dir string) (err error)
	TransformName(name string) (camel, class string)
	Replace(content, mod string) (tpl string)
	WriteToFile(fileName string, content string) error
}

type fileService struct {
}

func (s *fileService) IsExitsDir(dir string) (ok bool, err error) {
	if _, err = os.Stat(dir); err != nil {
		if !os.IsNotExist(err) {
			ok = true
			return
		}
	}
	return
}

func (s *fileService) IsExistsFile(path string) (ok bool, err error) {
	_, err = os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if !os.IsExist(err) {
			err = xerrors.Errorf("%w", err)
			return
		}

	}
	ok = true
	return
}

func (s *fileService) Mkdir(dir string) (err error) {
	var ok bool
	ok, err = s.IsExitsDir(dir)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	if ok {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

func (s *fileService) TransformName(name string) (camel, class string) {
	camel = util.CamelString(name)
	class = strings.Title(camel)
	return
}

func (s *fileService) Replace(content, mod string) (tpl string) {
	camel, class := s.TransformName(mod)
	tpl = strings.ReplaceAll(content, "juno", mod)
	tpl = strings.ReplaceAll(tpl, "greeting", camel)
	tpl = strings.ReplaceAll(tpl, "Greeting", class)
	return
}

func (s *fileService) WriteToFile(fileName string, content string) (err error) {
	var f *os.File
	f, err = os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	var n int64
	n, err = f.Seek(0, io.SeekEnd)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	_, err = f.WriteAt([]byte(content), n)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	return
}

func NewFileService(db *gorm.DB) FileService {
	return &fileService{}
}
