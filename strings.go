package strings

import (
	"regexp"
	"sort"
	"strings"
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

// CountWords return map with count of each words
func CountWords(s string) map[string]int {

	c := map[string]int{}
	words := strings.Split(s, " ")
	for _, w := range words {
		_, exist := c[w]
		if exist {
			c[w]++
		} else {
			c[w] = 1
		}
	}
	return c

}

// HighestOccurredWord returns most frequent word
func HighestOccurredWord(s string) string {
	countMap := CountWords(s)
	if len(countMap) == 0 {
		return ""
	}
	kv := Kvs{}
	for k, v := range countMap {
		kv = append(kv, KV{k, v})
	}
	sort.Sort(kv)
	return kv[len(kv)-1].Key
}

// RemoveAllNonAlphanumeric gets only nums and alphabetic symbols
// on string
func RemoveAllNonAlphanumeric(s string) (string, error) {

	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		return "", err
	}
	return reg.ReplaceAllString(s, ""), nil
}
