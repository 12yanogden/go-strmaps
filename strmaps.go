package strmaps

import (
	"fmt"
)

// Return a string representation of a string map
func Pretty(m map[string]string) string {
	out := ""

	for key, value := range m {
		out += key + ": " + value + "\n"
	}

	return out
}

// Print a string representation of a string map
func Print(m map[string]string) {
	fmt.Printf("%s", Pretty(m))
}
