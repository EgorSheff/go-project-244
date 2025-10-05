package parser

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSON(t *testing.T) {
	data := `{"a": 1, "b": "str"}`
	os.WriteFile("test.json", []byte(data), 0644)
	res, err := ParseFile("test.json")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, map[string]any{"a": 1.0, "b": "str"}, res)
	os.Remove("test.json")
}

func TestYAML(t *testing.T) {
	data := `a: 1
b: "str"`
	os.WriteFile("test.yaml", []byte(data), 0644)
	res, err := ParseFile("test.yaml")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, map[string]any{"a": 1, "b": "str"}, res)
	os.Remove("test.yaml")
}
