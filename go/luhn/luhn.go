package luhn

import (
	"fmt"
	"strconv"
	"strings"
)

func Valid( candidate string) bool {

	fmt.Println("candidate = ", candidate)

	// spaces should be stripped
	candidate = strings.ReplaceAll(candidate, " ", "")

	// strings of length 1 or less are invalid
	if len(candidate) <2 {
		return false
	}

	fmt.Println("stripped candidate:", candidate)

	// words containing non-digit chars are invalid
	// in other words: if candidate cannot be converted to number, its invalid
	_,err := strconv.Atoi(candidate)
	if err != nil {
		return false
	}

	// now, the algorithm itself
	length := len(candidate)

	for i:= 2; i<=length; i += 2 {
		c := candidate[length-i]

		fmt.Println("i ", i, ", c: ", string(c))
		val ,_ := strconv.Atoi(string(c))

		val = val * 2
		if val > 9 {
			val = val - 9
		}
		fmt.Println("val:", val)
		candidate = candidate[:length-i] + strconv.Itoa(val) + candidate[length-i+1:]
		fmt.Println("intermediate candidate: ", candidate)
	}

	sum := sumUpAllDigits( candidate )
	fmt.Println("sum of digits = ", sum)
	remainder := sum%10
	fmt.Println ("remainder % 10", remainder, "leaving funcion")
	return (remainder == 0)
}

func sumUpAllDigits(candidate string) int {
	sum := 0
	for _,c := range(candidate){
		n,_ := strconv.Atoi(string(c))
		sum += n
	}
	return sum
}