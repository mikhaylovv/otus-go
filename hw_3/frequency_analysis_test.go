package hw3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTop10SimplePositive(t *testing.T) {
	as := assert.New(t)
	text := `one one
two two two
three three three three
four four four four four
five five five five five five
six six six: six: six six six
seven seven seven: seven seven seven seven ;seven: 
eight eight, eight, eight, eight, eight eight eight eight
nine nine nine! nine! nine! nine! nine! nine nine nine
ten ten ten Ten Ten Ten Ten Ten ten ten ten
111 111 111 111 111 111 111 111 111 111 111 111 111 111 111 111 111 111 111 111 
here any other text for testing
`

	as.EqualValues(
		[]string{"ten", "nine", "eight", "seven", "six", "five", "four", "three", "two", "one"},
		Top10(text))
}

func TestTop10EmptyCheck(t *testing.T) {
	as := assert.New(t)
	as.EqualValues([]string{}, Top10(""))
}

func TestTop10SmallText(t *testing.T) {
	as := assert.New(t)
	as.EqualValues([]string{"small", "text"}, Top10("small small text :("))
}
