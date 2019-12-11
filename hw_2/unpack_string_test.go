package hw2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testFunc(as * assert.Assertions, in string, out string) {
	s, er := UnpackString(in)
	as.Nil(er)
	as.Equal(out, s)
}

func TestUnpackStringSimplePositive(t *testing.T) {
	as := assert.New(t)

	testFunc(as, "abcd", "abcd")
	testFunc(as, "a4bc2d5e", "aaaabccddddde")
}


func TestUnpackStringSimpleNegative(t *testing.T) {
	as := assert.New(t)

	s, er := UnpackString("45")
	as.Empty(s)
	as.NotNil(er)
}

func TestUnpackStringSimpleAdvanced(t *testing.T) {
	as := assert.New(t)

	testFunc(as, `qwe\4\5`, "qwe45")
	testFunc(as, `qwe\45`, "qwe44444")
	testFunc(as, `qwe\\5`, `qwe\\\\\`)
	testFunc(as, `\5qwe\\5`, `5qwe\\\\\`)
}