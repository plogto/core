package util

func PrepareKeyPattern(value string) string {
	return "$$$___" + value + "___$$$"
}
