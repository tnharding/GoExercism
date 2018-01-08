// Package acronym creates an acronym for the suppied string
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"strings"
)

// Abbreviate creates an acronym for the supplied string
func Abbreviate(s string) string {

	var acronym = ""

	//String is parsed based upon the result of calling
	//func Split
	words := strings.FieldsFunc(s, Split)
	for _, word := range words {
		acronym += strings.ToUpper(word[0:1])
	}
	return acronym
}

//Split determines if split is to occur with the supplied character
func Split(r rune) bool {
	return r == ' ' || r == '-'
}
