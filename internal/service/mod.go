package service

import (
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"golang.org/x/xerrors"
)

type ModService interface {
	GetMod() (mod string, err error)
}

type modService struct {
}

func (s *modService) GetMod() (mod string, err error) {
	file, err := os.Open("go.mod")
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	defer file.Close()
	var content []byte
	content, err = ioutil.ReadAll(file)
	reg, err := regexp.Compile(` \S+`)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	info := reg.Find(content)
	mod = strings.TrimSpace(string(info))
	return
}

func NewModService() ModService {
	return &modService{}
}
