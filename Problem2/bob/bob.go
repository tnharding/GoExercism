// Package bob receives a remark and returns a response.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
)

// Hey accepts a remark and will return a response
func Hey(remark string) string {

	//Lets remove all leading and trailing white space
	s := strings.TrimSpace(remark)
	if len(s) == 0 { // Bob didn't say anything to us
		return "Fine. Be that way!" // response when bob is quiet.
	}

	if s[len(s)-1:] == "?" { //you are asking me a question
		if s == strings.ToUpper(s) && strings.ToLower(s) != s { //are you yelling at me
			return "Calm down, I know what I'm doing!"
		}
		return "Sure."
	}

	//Don't yell at me
	if s == strings.ToUpper(s) && strings.ToLower(s) != s {
		return "Whoa, chill out!"
	}

	return "Whatever."
}
