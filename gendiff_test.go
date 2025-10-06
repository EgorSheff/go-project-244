package code

import (
	"code/parser"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenDiff(t *testing.T) {
	a, err := parser.ParseFile("testdata/file1.json")
	if err != nil {
		t.Fatal(err)
	}
	b, err := parser.ParseFile("testdata/file2.json")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, `{
  - follow: false
    host: hexlet.io
  - proxy: 123.234.53.22
  - timeout: 50
  + timeout: 20
  + verbose: true
}`, FormatDiffs(genDiff(a, b)))
}
