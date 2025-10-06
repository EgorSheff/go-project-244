package code

import (
	"fmt"
	"strings"
)

func FormatDiffs(difs []Diff) string {
	var b strings.Builder
	b.WriteString("{\n")
	for _, diff := range difs {
		b.WriteString("  ")
		if diff.Add {
			b.WriteString("+ ")
		}
		if diff.Del {
			b.WriteString("- ")
		}
		if !diff.Add && !diff.Del {
			b.WriteString("  ")
		}
		fmt.Fprintf(&b, "%s: %s\n", diff.Key, diff.Val)
	}
	b.WriteString("}")

	return b.String()
}
