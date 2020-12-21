package hamming

import (
	"errors"
)

func Distance(a, b string) (int, error) {

	// if different len, error
	if len(a) != len(b) {
		return -1, errors.New(("strings of different lengths are not allowed"))
	}

	if len(a) == 0 {
		return 0, nil
	}

	var distance int

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance, nil
}
