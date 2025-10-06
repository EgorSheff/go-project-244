package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSON(t *testing.T) {
	res, err := ParseFile("../testdata/file1.json")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, map[string]any{"host": "hexlet.io", "timeout": 50.0, "proxy": "123.234.53.22", "follow": false}, res)
}

func TestYAML(t *testing.T) {
	res, err := ParseFile("../testdata/file1.yml")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, map[string]any{"host": "hexlet.io", "timeout": 50.0, "proxy": "123.234.53.22", "follow": false}, res)
}

func TestDifferentFormats(t *testing.T) {
	jsn, err := ParseFile("../testdata/file1.json")
	if err != nil {
		t.Fatal(err)
	}
	yml, err := ParseFile("../testdata/file1.yml")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, jsn, yml)
}
