//Package twofer is short for two for one. One for you and one for me.
package twofer

import "fmt"

//ShareWith return sentence "One for X, one for me.". X is given name or "you" if name is not provided.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
