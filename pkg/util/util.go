package util

import (
	"golang.org/x/xerrors"
)

func Unwrap(err error) (uerr error) {
	for err != nil {
		uerr = err
		err = xerrors.Unwrap(err)

	}
	return
}
