package version

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
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

func init() {
	srcDir := fmt.Sprintf("%s/pkg/mod/github.com/go-juno/", "/Users/joker/go")
	fileList, err := ioutil.ReadDir(srcDir)
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
