package service

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/go-juno/juno/pkg/util"
	"golang.org/x/xerrors"
)

type FileService interface {
	IsExitsDir(dir string) (ok bool, err error)
	IsExistsFile(path string) (ok bool, err error)
	Mkdir(dir string) (err error)
	TransformName(name string) (camel, class string)
	Replace(content, mod string) (tpl string)
	WriteToFile(fileName string, content string) error
	CopyPath(src, dst string) (ok bool, err error)
	CopyFile(src, dst string) (ok bool, err error)
	ReadAll(filePth string) ([]byte, error)
	ReplaceAll(root, old, new string) (err error)
	ReplaceMod(root string) error
	ReplaceMain(root, old, new string) error
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

func (s *fileService) GetFileInfo(src string) os.FileInfo {
	if fileInfo, e := os.Stat(src); e != nil {
		if os.IsNotExist(e) {
			return nil
		}
		return nil
	} else {
		return fileInfo
	}
}

func (s *fileService) CopyFile(src, dst string) (ok bool, err error) {
	if len(src) == 0 || len(dst) == 0 {
		return
	}
	src = strings.Replace(src, "\\", "/", -1)
	srcFile, err := os.OpenFile(src, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return
	}
	defer srcFile.Close()

	dst = strings.Replace(dst, "\\", "/", -1)
	dstPathArr := strings.Split(dst, "/")
	dstPathArr = dstPathArr[0 : len(dstPathArr)-1]
	dstPath := strings.Join(dstPathArr, "/")
	dstFileInfo := s.GetFileInfo(dstPath)
	if dstFileInfo == nil {
		if err = os.MkdirAll(dstPath, os.ModePerm); err != nil {
			return
		}
	}

	//这里要把O_TRUNC 加上，否则会出现新旧文件内容出现重叠现象
	dstFile, err := os.OpenFile(dst, os.O_CREATE|os.O_TRUNC|os.O_RDWR, os.ModePerm)
	if err != nil {
		return
	}
	defer dstFile.Close()

	if _, err = io.Copy(dstFile, srcFile); err != nil {
		return
	}
	ok = true
	return
}

func (s *fileService) ReadAll(filePth string) ([]byte, error) {
	f, err := os.Open(filePth)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}

func (s *fileService) CopyPath(src, dst string) (ok bool, err error) {
	src = strings.Replace(src, "\\", "/", -1)
	srcFileInfo := s.GetFileInfo(src)
	if srcFileInfo == nil || !srcFileInfo.IsDir() {
		return
	}

	err = filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			err = xerrors.Errorf("%w", err)
			return err
		}

		path = strings.Replace(path, "\\", "/", -1)
		relationPath := strings.Replace(path, src, "", -1)
		dstPath := strings.TrimRight(strings.TrimRight(strings.Replace(dst, "\\", "/", -1), "/"), "\\") + relationPath

		if !info.IsDir() {
			var ok bool
			ok, err = s.CopyFile(path, dstPath)
			if err != nil {
				err = xerrors.Errorf("%w", err)
				return err
			}
			if ok {
				return nil
			} else {
				return errors.New(path + " copy fail")
			}
		} else {
			if _, err := os.Stat(dstPath); err != nil {
				if os.IsNotExist(err) {
					if err := os.MkdirAll(dstPath, os.ModePerm); err != nil {
						return err
					} else {
						return nil
					}
				} else {
					return err
				}
			} else {
				return nil
			}
		}
	})
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	ok = true
	return
}

func (s *fileService) ReplaceAll(root, old, new string) (err error) {
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			// 替换内容
			var text []byte
			text, err = s.ReadAll(path)

			if err != nil {
				err = xerrors.Errorf("%w", err)
				return err
			}
			str := string(text)
			str = strings.ReplaceAll(str, old, new)
			if err := s.WriteToFile(path, str); err != nil {
				return err
			}
		}

		return err
	})
	return err
}

func (s *fileService) ReplaceMod(root string) error {
	path := fmt.Sprintf("%s/go.mod", root)
	text, err := s.ReadAll(path)
	if err != nil {
		return err
	}
	str := string(text)
	reg := regexp.MustCompile(`(replace \([\s\S]*?\))`)
	str = reg.ReplaceAllString(str, "")
	if err := s.WriteToFile(path, str); err != nil {
		return err
	}
	return nil
}

func (s *fileService) ReplaceMain(root, old, new string) error {
	path := fmt.Sprintf("%s/main.go", root)
	text, err := s.ReadAll(path)
	if err != nil {
		return err
	}
	str := string(text)
	reg := regexp.MustCompile(old)
	str = reg.ReplaceAllString(str, new)
	if err := s.WriteToFile(path, str); err != nil {
		return err
	}
	return nil
}

func NewFileService() FileService {
	return &fileService{}
}
