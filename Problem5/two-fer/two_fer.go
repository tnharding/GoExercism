// Two-fer or 2-fer is short for two for one. One for you and one for me.
// "One for X, one for me."
// Where X is a name or "you".
// If the given name is "Alice", the result should be "One for Alice, one for me."
// If no name is given, the result should be "One for you, one for me."

package twofer

import (
	"fmt"
)

// ShareWith will return a response specific to a supplied name.
func ShareWith(name string) string {

	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
