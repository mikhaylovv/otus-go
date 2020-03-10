package hw7

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestReadDir(t *testing.T) {
	const dirName = "test_dir"
	assert.NoError(t, os.Mkdir(dirName, os.ModePerm))
	defer func() { _ = os.RemoveAll(dirName) }()

	assert.NoError(t, ioutil.WriteFile("test_dir/VAR_TEST_134", []byte("134"), os.ModePerm))
	assert.NoError(t, ioutil.WriteFile("test_dir/VAR_TEST_888", []byte("888"), os.ModePerm))

	m, err := ReadDir(dirName)

	assert.NoError(t, err)
	assert.Equal(t, 2, len(m))
	assert.Equal(t, "134", m["VAR_TEST_134"])
	assert.Equal(t, "888", m["VAR_TEST_888"])
}
