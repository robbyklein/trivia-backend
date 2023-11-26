package helpers

import "unicode/utf8"

func IsPasswordValid(password string) bool {
	length := utf8.RuneCountInString(password)
	return length >= 8
}
