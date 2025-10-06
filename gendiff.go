package code

import (
	"code/parser"
	"fmt"
	"reflect"
	"slices"
	"sort"
)

type Diff struct {
	Add bool
	Del bool
	Key string
	Val any
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

	if format == "" || format == "stylish" {
		return FormatDiffsStylish(diffs, 0), nil
	}
	return "", ErrUnsupportedFormatter
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

		nestA, aNestOk := aV.(map[string]any)
		nestB, bNestOk := bV.(map[string]any)

		if aNestOk && bNestOk && reflect.DeepEqual(nestA, nestB) {
			res = append(res, Diff{
				Key: key,
				Val: parseValue(aV),
			})
			continue
		}

		if aNestOk && bNestOk {
			res = append(res, Diff{
				Key: key,
				Val: genDiff(nestA, nestB),
			})
			continue
		}

		if aOk && bOk && aV == bV {
			res = append(res, Diff{
				Key: key,
				Val: parseValue(aV),
			})
			continue
		}

		if aOk {
			res = append(res, Diff{
				Del: true,
				Key: key,
				Val: parseValue(aV),
			})
		}
		if bOk {
			res = append(res, Diff{
				Add: true,
				Key: key,
				Val: parseValue(bV),
			})
		}
	}

	return res
}

func parseValue(value any) any {
	if m, ok := value.(map[string]any); ok {
		res := make([]Diff, 0, len(m))
		for k, v := range m {
			res = append(res, Diff{
				Key: k,
				Val: parseValue(v),
			})
		}

		slices.SortFunc(res, func(a Diff, b Diff) int {
			if a.Key >= b.Key {
				return 1
			}
			return -1
		})
		return res
	}
	return value
}
