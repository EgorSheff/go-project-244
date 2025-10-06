package code

import (
	"code/parser"
	"fmt"
	"testing"
)

func TestGenDiff(t *testing.T) {
	a, err := parser.ParseFile("test/file1.json")
	if err != nil {
		t.Fatal(err)
	}
	b, err := parser.ParseFile("test/file2.json")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(genDiff(a, b))
}
