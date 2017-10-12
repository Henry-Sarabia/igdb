package igdb

import (
	"strconv"
	"strings"
)

// URL is
type URL string

// Tag codes??
type Tag int //uint32

// intsToStrings returns the slice of strings
// equivalent to the list of ints.
func intsToStrings(ints []int) []string {
	var str []string
	for _, i := range ints {
		str = append(str, strconv.Itoa(i))
	}
	return str
}

// intsToCommaString returns a comma separated
// list of ints as a single string.
func intsToCommaString(ints []int) string {
	s := intsToStrings(ints)
	return strings.Join(s, ",")
}
