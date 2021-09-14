package util

import (
	"strings"

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
