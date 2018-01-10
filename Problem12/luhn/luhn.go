//Package luhn ...
//Given a number determine whether or not it is valid per the Luhn formula.
package luhn

import (
	"strconv"
	"strings"
)

//Replace space chars with nothing
var replacer = strings.NewReplacer(" ", "")

func Valid(num string) bool {

	//Remove all space chars from num
	s := replacer.Replace(num)

	//check to see if we have any non digit characters in string
	if _, err := strconv.Atoi(s); err != nil {
		return false
	}

	//length of 1 or less is invalid
	if len(s) < 2 {
		return false
	}

	numInts := convertNumToIntSlice(s)

	//double every second digit
	//if value > 9 subtract nine from the value
	for i := len(numInts) - 2; i >= 0; i = i - 2 {
		numInts[i] = numInts[i] * 2
		if numInts[i] > 9 {
			numInts[i] = numInts[i] - 9
		}
	}

	//Add all the digits together
	var total int
	for i := 0; i < len(numInts); i++ {
		total += numInts[i]
	}

	//if total is not a factor of 10
	//return invalid number
	if total%10 != 0 {
		return false
	}

	//number is valid
	return true
}

func convertNumToIntSlice(num string) []int {
	numSlice := make([]int, len(num))

	//convert rune to int value
	for i, v := range num {
		numSlice[i] = int(v - 0x30)
	}
	return numSlice
}
