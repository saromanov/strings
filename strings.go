package strings

// IsContainsNumbers return true if input string contains numbers
func IsContainsNumbers(s string) bool {
	
	for _, i := range s {
		if i >= '0' && i < '9' {
			return true
		}
	}
	return false
}