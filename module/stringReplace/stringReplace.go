package stringReplace

import "strings"

func RemoveEnd(target string) string {
	target = strings.Replace(target, "\n", "", -1)
	target = strings.Replace(target, "\r", "", -1)
	return target
}
