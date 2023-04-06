package util

import "github.com/google/uuid"

func PrepareKeyPattern(value uuid.UUID) string {
	return "$$$___" + value.String() + "___$$$"
}
