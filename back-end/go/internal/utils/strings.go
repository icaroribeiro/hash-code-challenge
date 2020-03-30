package utils

import (
	"strings"
)

// It removes escape sequences from a string.
func RemoveEscapeSequencesFromString(s string, sequences ...string) string {
	var sequence string

	for _, sequence = range sequences {
		s = strings.Replace(s, sequence, "", -1)
	}

	return s
}
