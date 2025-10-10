package code

import (
	"code/formatters"
	"code/getdiff"
	"code/parser"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStylishDiff(t *testing.T) {
	a, err := parser.ParseFile("testdata/file1.json")
	if err != nil {
		t.Fatal(err)
	}
	b, err := parser.ParseFile("testdata/file2.json")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, `{
  common: {
    + follow: false
      setting1: Value 1
    - setting2: 200
    - setting3: true
    + setting3: null
    + setting4: blah blah
    + setting5: {
          key5: value5
      }
      setting6: {
          key: value
        + ops: vops
      }
  }
  group1: {
    - baz: bas
    + baz: bars
      foo: bar
    - nest: {
          key: value
      }
    + nest: str
  }
- group2: {
      abc: 12345
      deep: {
          id: 45
      }
  }
+ group3: {
      deep: {
          id: {
              number: 45
          }
      }
      fee: 100500
  }
}`, formatters.FormatDiffsStylish(getdiff.GetDiffs(a, b)))
}

func TestPlainDiff(t *testing.T) {
	a, err := parser.ParseFile("testdata/file1.json")
	if err != nil {
		t.Fatal(err)
	}
	b, err := parser.ParseFile("testdata/file2.json")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, `Property 'common.follow' was added with value: false
Property 'common.setting2' was removed
Property 'common.setting3' was updated. From true to null
Property 'common.setting4' was added with value: 'blah blah'
Property 'common.setting5' was added with value: [complex value]
Property 'common.setting6.ops' was added with value: 'vops'
Property 'group1.baz' was updated. From 'bas' to 'bars'
Property 'group1.nest' was updated. From [complex value] to 'str'
Property 'group2' was removed
Property 'group3' was added with value: [complex value]`, formatters.FormatDiffsPlain(getdiff.GetDiffs(a, b)))
}
