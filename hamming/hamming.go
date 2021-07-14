package hamming

import (
	"errors"
)

// Distance returns the hamming distance between two DNA strands
func Distance(a, b string) (distance int, err error) {

	if len(a) != len(b) {
		err = errors.New("")

		return
	}

	for index, char := range a {
		if char != rune(b[index]) {
			distance++
		}
	}

	return
}
