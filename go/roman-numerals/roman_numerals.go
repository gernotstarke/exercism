package romannumerals

import (
	"errors"
	"fmt"
)

var values = [...]int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

var syms = [...]string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

func ToRomanNumeral(arabic int) (string, error) {

	fmt.Printf("arabic = %d", arabic)

	if arabic <= 0 {
		return "", errors.New("can't work with zero or negative numbers")
	}

	if arabic > 3000 {
		return "", errors.New("cannot convert numbers greater than 3000")
	}

	roman := ""

	var i int = 0

	for arabic > 0 {
		current := arabic / values[i]
		for j := 1; j <= current; j++ {
			roman = roman + syms[i]
		}
		arabic = arabic - current*values[i]
		i++
	}

	fmt.Println(", roman = ", roman)
	return roman, nil
}
