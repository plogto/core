package util

import (
	"strings"
)

func IsImage(contentType string) bool {
	return strings.Split(contentType, "/")[0] == "image"
}
