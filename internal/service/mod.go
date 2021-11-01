package service

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/go-juno/juno/pkg/util"
	"golang.org/x/xerrors"
)

type ModService interface {
	GetMod() (mod string, err error)
}

type modService struct {
}

func (s *modService) GetMod() (mod string, err error) {
	modFile := filepath.Join(util.GetPwd(), "go.mod")
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	file, err := os.Open(modFile)
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
