package code

import (
	"code/parser"
	"fmt"
	"sort"
)

type Diff struct {
	Add bool
	Del bool
	Key string
	Val string
}

func GenDiff(filepathA, filepathB string, format string) (string, error) {
	a, err := parser.ParseFile(filepathA)
	if err != nil {
		return "", fmt.Errorf("error parse file %s: %v", filepathA, err)
	}
	b, err := parser.ParseFile(filepathB)
	if err != nil {
		return "", fmt.Errorf("error parse file %s: %v", filepathB, err)
	}

	diffs := genDiff(a, b)
	return FormatDiffs(diffs), nil
}

func genDiff(a map[string]any, b map[string]any) []Diff {
	var keys []string
	for k := range a {
		keys = append(keys, k)
	}
	for k := range b {
		if _, ok := a[k]; !ok {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)

	var res []Diff
	for _, key := range keys {
		aV, aOk := a[key]
		bV, bOk := b[key]

		if aOk && bOk && aV == bV {
			res = append(res, Diff{
				Key: key,
				Val: fmt.Sprint(aV),
			})
			continue
		}
		if aOk {
			res = append(res, Diff{
				Del: true,
				Key: key,
				Val: fmt.Sprint(aV),
			})
		}
		if bOk {
			res = append(res, Diff{
				Add: true,
				Key: key,
				Val: fmt.Sprint(bV),
			})
		}
	}

	return res
}
