//Package hamming ...
//The Hamming distance is found by comparing two DNA strands and
//counting how many of the nucleotides are different from their equivalent
//in the other string.
//     GAGCCTACTAACGGGAT
//     CATCGTAATGACGGCCT
//     ^ ^ ^  ^ ^    ^^
// The Hamming distance between these two DNA strands is 7.
// The Hamming distance is only defined for sequences of equal length.
package hamming

import (
	"errors"
)

//Distance calculate the Hamming difference between two DNA strands.
func Distance(a, b string) (int, error) {
	var distance = 0

	//If DNA sequences are not of equal length return error
	if len(a) != len(b) {
		err := errors.New("sequences not of equal length")
		return -1, err
	}

	//Count the differences in each strand to the other strand
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance, nil
}
