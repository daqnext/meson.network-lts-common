package mesonUtils

import (
	"strings"
)

//VersionCompare
//compare version (x.x.x)
//
//a>b return 1
//
//a==b return 0
//
//a<b  return -1
func VersionCompare(a string, b string) int {
	aVersion := strings.Split(a, ".")
	bVersion := strings.Split(b, ".")
	for i, _ := range aVersion {
		if aVersion[i] > bVersion[i] {
			return 1
		} else if aVersion[i] < bVersion[i] {
			return -1
		}
	}
	return 0
}
