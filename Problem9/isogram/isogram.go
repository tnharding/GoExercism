//Package isogram ...
// Determine if a word or phrase is an isogram.
// An isogram (also known as a "nonpattern word") is a word or phrase without a repeating letter.
// Examples of isograms:
// - lumberjacks
// - background
// - downstream
// The word *isograms*, however, is not an isogram, because the s repeats.
package isogram

import (
	"strings"
	"unicode"
)

//IsIsogram determines is a phrase is an isogram.
func IsIsogram(isogram string) bool {

	//the map will hold an instance of a letter
	isomap := make(map[rune]int)

	//convert string to lower case so an 'A' and 'a' will
	//be seen as a non isogram.
	s := strings.ToLower(isogram)

	//scan the supplied phrase
	for _, v := range s {
		//check to see if the character is already in the map
		if _, ok := isomap[v]; ok {
			return false
		}
		//Only add letters to the map
		if unicode.IsLetter(v) {
			isomap[v] = 1
		}
	}

	return true
}
