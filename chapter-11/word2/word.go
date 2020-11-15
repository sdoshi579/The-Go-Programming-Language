package word

import (
	"unicode"
)

// IsPalindrome reports whether s is palindrom string or not
func IsPalindrome(s string) bool {
	var letters []rune
	for _, i := range s {
		if unicode.IsLetter(i) {
			letters = append(letters, unicode.ToLower(i))
		}
	}
	for i := range letters {
		if letters[i] != letters[len(letters)-i-1] {
			return false
		}
	}
	return true
}
