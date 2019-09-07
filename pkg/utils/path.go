package utils

import (
	"path"
	"strings"
)

func RemoveDirFromPath(input string, index int) string {
	sliced := strings.Split(input[1:], "/")
	if index >= len(sliced) {
		return input
	}
	trimmed := append(sliced[:index], sliced[index+1:]...)
	return "/" + path.Join(trimmed...)
}
