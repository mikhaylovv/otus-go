package hw3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFrequencyAnalysisSimplePositive(t *testing.T) {
	as := assert.New(t)
	text := `one one
two two two
three three three three
four four four four four
five five five five five five
six six six six six six six
seven seven seven seven seven seven seven seven 
eight eight eight eight eight eight eight eight eight
nine nine nine nine nine nine nine nine nine nine
ten ten ten ten ten ten ten ten ten ten ten
here any other text for testing
`

	res := FrequencyAnalysis(text)
	expect := [10]string{"ten", "nine", "eight", "seven", "six", "five", "four", "three", "two", "one"}

	as.EqualValues(expect, res)
}
