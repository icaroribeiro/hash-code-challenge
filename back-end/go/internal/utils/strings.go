package utils

import (
	"strings"
)

// It removes escape sequences from a string.
func RemoveEscapeSequences(s string, sequences ...string) string {
	var i int

	for i = range sequences {
		s = strings.Replace(s, sequences[i], "", -1)
	}

	return s
}
