package parser

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSON(t *testing.T) {
	res, err := ParseFile("../testdata/file1.json")
	if err != nil {
		t.Fatal(err)
	}

	d, err := os.ReadFile("../testdata/file1.json")
	if err != nil {
		t.Fatal(err)
	}
	var expected map[string]any
	if err := json.Unmarshal(d, &expected); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, res)
}

func TestYAML(t *testing.T) {
	res, err := ParseFile("../testdata/file1.yml")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, map[string]any{"host": "hexlet.io", "timeout": 50.0, "proxy": "123.234.53.22", "follow": false}, res)
}
