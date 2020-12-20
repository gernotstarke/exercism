// acronym exercise from exercism.io
// author: Gernot Starke
package acronym

import (
	"strings"
)

func firstLetter(s string) string {
	plain := strings.TrimSpace(strings.Trim(s, "_-!.,;: "))

	return string(plain[0])
}

// Abbreviate joins and uppercases the first letters, e.g. "Portable Network Graphics" -> PNG.
func Abbreviate(s string) string {
	result := ""

	for _, v := range strings.Fields(strings.Replace(s, "-", " ", -1)) {
		//fmt.Println(i, result, v)
		result = result + strings.ToUpper(firstLetter(v))
	}

	return result
}
