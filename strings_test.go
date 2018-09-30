package strings

import (
	"fmt"
	"testing"
)

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

func TestHighestOccurredWord(t *testing.T) {
	word := HighestOccurredWord("hello world world foobar world")
	if word != "world" {
		t.Errorf("unable to get highest occurred word")
	}
}

func TestRemoveAllNonAlphanumeric(t *testing.T) {
	s, err := RemoveAllNonAlphanumeric("Yeah!")
	if err != nil {
		t.Errorf(fmt.Sprintf("%v", err))
	}

	if s != "Yeah" {
		t.Errorf("unable to remove non alphanumeric symbols")
	}

	s2, err := RemoveAllNonAlphanumeric("Yeah")
	if err != nil {
		t.Errorf("unable to remove non alphanumeric symbols")
	}
	if s2 != "Yeah" {
		t.Errorf("unable to remove non alphanumeric symbols")
	}
}

func TestValidateEmail(t *testing.T) {
	if !ValidateEmail("123@mail.ru") {
		t.Errorf("unable to validate email")
	}
	if ValidateEmail("123@@mail.ru") {
		t.Errorf("unable to validate email")
	}
}
