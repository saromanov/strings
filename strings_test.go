package strings

import "testing"

func TestIsContainsNumbers(t *testing.T) {
	if IsContainsNumbers("abc") {
		t.Errorf("string not contains numbers")
	}
	if !IsContainsNumbers("1abc") {
		t.Errorf("string contains numbers")
	}
	if !IsContainsNumbers("0") {
		t.Errorf("string contains numbers")
	}
	if !IsContainsNumbers("&^uj.5sss") {
		t.Errorf("string contains numbers")
	}
}
