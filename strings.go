package strings

import (
	"regexp"
)

// IsContainsNumbers return true if input string contains numbers
func IsContainsNumbers(s string) bool {

	return IsContainsFunc(s, func(str string) bool {
		for _, i := range s {
			if i >= '0' && i < '9' {
				return true
			}
		}
		return false
	})
}

// IsContainsMap returns true if defined method
// contains target string
func IsContainsFunc(s string, f func(string) bool) bool {
	return f(s)
}

// ExtractNumbers returns slice of string with numbers if this is exist
func ExtractNumbers(s string) []string {
	re := regexp.MustCompile("[0-9]+")
	return re.FindAllString(s, -1)
}
