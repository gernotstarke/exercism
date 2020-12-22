package reverse

import "strings"

func Reverse( orig string) string {
	if orig == "" {
		return ""
	}

	chars := strings.Split( orig,"")

	var result string = ""
	for i:=len(chars)-1;i>=0;i-- {
		result += chars[i]
	}
	return result
}
