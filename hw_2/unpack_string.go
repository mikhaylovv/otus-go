package hw2

import (
	"errors"
	"strings"
	"unicode"
	"unicode/utf8"
)

/*UnpackString ...
	Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:

	* "a4bc2d5e" => "aaaabccddddde"
	* "abcd" => "abcd"
	* "45" => "" (некорректная строка)

	Дополнительное задание: поддержка escape - последовательности
	* `qwe\4\5` => `qwe45` (*)
	* `qwe\45` => `qwe44444` (*)
	* `qwe\\5` => `qwe\\\\\` (*)
 */
func UnpackString(str string) (string, error) {

	var builder strings.Builder

	prevLetter, size := utf8.DecodeRuneInString(str)
	if prevLetter == utf8.RuneError {
		if size == 0 {
			return "", errors.New("empty string passed")
		}
		return "", errors.New("encoding is invalid")
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
