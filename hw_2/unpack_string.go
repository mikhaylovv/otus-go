package hw_2

import (
	"errors"
	"strings"
	"unicode"
	"unicode/utf8"
)

func UnpackString(str string) (string, error) {

	var builder strings.Builder

	prevLetter, size := utf8.DecodeRuneInString(str)
	if prevLetter == utf8.RuneError {
		if size == 0 {
			return "", errors.New("empty string passed")
		} else {
			return "", errors.New("encoding is invalid")
		}
	}

	if unicode.IsDigit(prevLetter) {
		return "", errors.New("first letter is digit")
	}

	builder.WriteRune(prevLetter)

	for i, r := range str {
		if i == 0 {
			continue
		}

		if unicode.IsDigit(r) {
			for i := 0; i < int(r-'0')-1; i++ {
				builder.WriteRune(prevLetter)
			}
			continue
		}

		prevLetter = r
		builder.WriteString(string(prevLetter))
	}

	return builder.String(), nil
}
