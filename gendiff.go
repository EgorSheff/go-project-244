package code

import (
	"code/formatters"
	"code/getdiff"
	"code/parser"
	"fmt"
)

func GenDiff(filepathA, filepathB string, format string) (string, error) {
	a, err := parser.ParseFile(filepathA)
	if err != nil {
		return "", fmt.Errorf("error parse file %s: %v", filepathA, err)
	}
	b, err := parser.ParseFile(filepathB)
	if err != nil {
		return "", fmt.Errorf("error parse file %s: %v", filepathB, err)
	}

	diffs := getdiff.GetDiffs(a, b)

	switch format {
	case "":
		fallthrough
	case "stylish":
		return formatters.FormatDiffsStylish(diffs), nil
	case "plain":
		return formatters.FormatDiffsPlain(diffs), nil
	}
	return "", formatters.ErrUnsupportedFormatter
}
