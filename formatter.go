package code

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrUnsupportedFormatter = errors.New("unsupported formater")
)

func FormatDiffsStylish(difs []Diff, level int) string {
	var b strings.Builder
	intent := strings.Repeat("  ", level*2)
	b.WriteString("{\n")
	for _, diff := range difs {
		b.WriteString(intent)
		if diff.Add {
			b.WriteString("+ ")
		}
		if diff.Del {
			b.WriteString("- ")
		}
		if !diff.Add && !diff.Del {
			b.WriteString("  ")
		}

		if nestDiffs, ok := diff.Val.([]Diff); ok {
			fmt.Fprintf(&b, "%s: %s\n", diff.Key, FormatDiffsStylish(nestDiffs, level+1))
			continue
		}
		if diff.Val == nil {
			fmt.Fprintf(&b, "%s: null\n", diff.Key)
			continue
		}
		fmt.Fprintf(&b, "%s: %s\n", diff.Key, fmt.Sprint(diff.Val))
	}
	b.WriteString(strings.Repeat("  ", max(2*level-1, 0)) + "}")

	return b.String()
}
