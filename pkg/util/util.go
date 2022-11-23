package util

import (
	"os"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"golang.org/x/tools/go/packages"
	"golang.org/x/xerrors"
)

func Unwrap(err error) (uerr error) {
	for err != nil {
		uerr = err
		err = xerrors.Unwrap(err)

	}
	return
}

func SnakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	return strings.ToLower(string(data[:]))
}

func TitleString(s string) string {
	caser := cases.Title(language.BrazilianPortuguese)
	return caser.String(s)
}

func CamelString(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := true
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if i == 0 && d >= 'A' && d <= 'Z' {
			d = d + 32
		}
		if !k && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || !k) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}

func GetPwd() (dir string) {
	dir, _ = os.Getwd()
	return
}

func GetMod() (mod string, err error) {
	cfg := &packages.Config{
		Mode:  packages.NeedFiles,
		Tests: false,
	}
	pkgs, err := packages.Load(cfg)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	mod = pkgs[0].String()
	return
}

func GetMethod(name string) (method string) {
	method = "POST"
	if strings.HasPrefix(name, "get") {
		method = "GET"
	} else if strings.HasPrefix(name, "create") {
		method = "POST"
	} else if strings.HasPrefix(name, "update") {
		method = "PUT"
	} else if strings.HasPrefix(name, "delete") || strings.HasPrefix(name, "cancel") {
		method = "DELETE"
	}
	return

}
