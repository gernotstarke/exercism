package isogram

import "strings"

func IsIsogram(word string) bool {
	stats := make(map[string]int)

	// remove "-" and " " from word
	word = strings.ReplaceAll(word, "-", "")
	word = strings.ReplaceAll(word, " ", "")

	iso := true
	for _, c := range word {
		upC := strings.ToUpper(string(c))
		stats[string(upC)]++
		if stats[string(upC)] > 1 {
			iso = false
		}
	}

	return iso
}
