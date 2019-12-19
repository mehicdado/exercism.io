/*
Package hamming counts Hmming distance between two strands of DNA
*/
package hamming

import "errors"

/*
Distance count difference between two DNA strands represented  by strings.
This is called "Hamming Distance". Strands must have the same lenght.
*/
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("Strands must be same lenght")
	}

	if a == b {
		return 0, nil
	}

	//count num of differences between two strands
	numOfDiff := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			numOfDiff++
		}
	}

	return numOfDiff, nil
}
