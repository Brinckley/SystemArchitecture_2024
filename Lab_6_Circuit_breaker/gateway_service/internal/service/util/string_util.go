package util

import "strings"

func ConvertToStringWithoutBlanks(bytes []byte) string {
	return RemoveAllBlanks(string(bytes))
}

func RemoveAllBlanks(tmp string) string {
	raw := strings.Replace(tmp, " ", "", -1)
	raw = strings.Replace(raw, "\n", "", -1)
	return raw
}
