package version

import (
	"io/ioutil"
	"log"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"golang.org/x/xerrors"
)

var Version = "v1.0.0"

func CompareVersion(version1 string, version2 string) int {
	version1 = strings.ReplaceAll(version1, "v", "")
	version2 = strings.ReplaceAll(version2, "v", "")
	versionA := strings.Split(version1, ".")
	versionB := strings.Split(version2, ".")

	for i := len(versionA); i < 4; i++ {
		versionA = append(versionA, "0")
	}
	for i := len(versionB); i < 4; i++ {
		versionB = append(versionB, "0")
	}
	for i := 0; i < 4; i++ {
		version1, _ := strconv.Atoi(versionA[i])
		version2, _ := strconv.Atoi(versionB[i])
		if version1 == version2 {
			continue
		} else if version1 > version2 {
			return 1
		} else {
			return -1
		}
	}
	return 0
}

func GetGoEnvPath() (path string, err error) {
	envCmd := "go env GOPATH"
	cmd := exec.Command("go", "env", "GOPATH")
	out, err := cmd.CombinedOutput()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	path = string(out[:len(out)-1])
	if path == "" {
		err = xerrors.Errorf("$GOPATH is not configured, see '%s'\n", envCmd)
		return
	}
	return
}

func init() {
	path, err := GetGoEnvPath()
	if err != nil {
		panic(err)
	}
	fileList, err := ioutil.ReadDir(filepath.Join(path, "pkg/mod/github.com/go-juno/"))
	if err != nil {
		log.Println("err", err)
	}
	for _, file := range fileList {
		list := strings.Split(file.Name(), "juno@")
		if len(list) == 2 {
			if CompareVersion(Version, list[1]) == -1 {
				Version = list[1]
			}
		}

	}
}
