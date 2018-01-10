//Package raindrops ...
//Convert a number to a string, the contents of which
//depend on the number's factors.
package raindrops

import "strconv"

// Convert a number to a string
func Convert(number int) string {
	var result string

	// - If the number has 3 as a factor, output 'Pling'.
	if number%3 == 0 {
		result += "Pling"
	}

	// - If the number has 5 as a factor, output 'Plang'.
	if number%5 == 0 {
		result += "Plang"
	}

	// - If the number has 7 as a factor, output 'Plong'.
	if number%7 == 0 {
		result += "Plong"
	}

	//   just pass the number's digits straight through.
	if result == "" {
		result = strconv.Itoa(number)
	}
	return result
}
