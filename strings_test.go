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

func TestIsContainsFunc(t *testing.T) {
	containsTest := IsContainsFunc("abc", func(str string) bool {
		for i, c := range str {
			if i == 1 && c == 'b' {
				return true
			}
		}
		return false
	})
	if !containsTest {
		t.Errorf("unable to apply IsContainsFunc")
	}
}

func TestExtractNumbers(t *testing.T) {
	if len(ExtractNumbers("254")) == 0 {
		t.Errorf("unable to extract numbers")
	}
}

func TestCountWords(t *testing.T) {
	countMap := CountWords("hello world")
	w, ok := countMap["world"]
	if !ok {
		t.Errorf("unable to count words")
	}
	if w != 1 {
		t.Errorf("unable to count words")
	}
}
