package getdiff

import (
	"reflect"
	"sort"
)

type Diff struct {
	Key       string
	Del       bool
	Add       bool
	OldVal    any
	NewVal    any
	NestDiffs []Diff
}

func (d Diff) Equal() bool {
	return d.Add && d.Del && reflect.DeepEqual(d.OldVal, d.NewVal)
}

func GetDiffs(a map[string]any, b map[string]any) []Diff {
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

		diff := Diff{
			Key:    key,
			Del:    aOk,
			Add:    bOk,
			OldVal: aV,
			NewVal: bV,
		}

		if aNestOk && bNestOk && !reflect.DeepEqual(nestA, nestB) {
			diff.NestDiffs = GetDiffs(nestA, nestB)
		}

		res = append(res, diff)
	}

	return res
}
