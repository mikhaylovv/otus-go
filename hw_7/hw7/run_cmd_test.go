package hw7

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"log"
	"strings"
	"testing"
)

func TestRunCmd(t *testing.T) {
	m := make(map[string]string, 2)
	m["VAR_TEST_123"] = "123"
	m["VAR_TEST_888"] = "888"

	c := []string{"printenv"}

	var buf bytes.Buffer
	log.SetOutput(&buf)

	err := RunCmd(c, m)

	assert.NoError(t, err)
	env := buf.String()
	assert.True(t, true, strings.Contains(env, "VAR_TEST_123=123"))
	assert.True(t, true, strings.Contains(env, "VAR_TEST_888=888"))
}
