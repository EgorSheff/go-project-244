package formatters

import (
	"code/getdiff"
	"errors"
	"fmt"
	"maps"
	"slices"
	"strings"
)

var (
	ErrUnsupportedFormatter = errors.New("unsupported formater")
)

func FormatDiffsStylish(diffs []getdiff.Diff) string {
	return formatDiffsStylish(diffs, 0)
}

func formatDiffsStylish(diffs []getdiff.Diff, level int) string {
	var b strings.Builder
	intent := strings.Repeat("  ", level*2)
	b.WriteString("{\n")
	for _, diff := range diffs {
		if len(diff.NestDiffs) > 0 {
			b.WriteString(intent)
			fmt.Fprintf(&b, "  %s: %s\n", diff.Key, formatDiffsStylish(diff.NestDiffs, level+1))
			continue
		}

		if diff.Equal() {
			b.WriteString(intent)
			fmt.Fprintf(&b, "  %s: %s\n", diff.Key, formatValStylish(diff.OldVal, level+1))
			continue
		}

		if diff.Del {
			b.WriteString(intent)
			fmt.Fprintf(&b, "- %s: %s\n", diff.Key, formatValStylish(diff.OldVal, level+1))
		}

		if diff.Add {
			b.WriteString(intent)
			fmt.Fprintf(&b, "+ %s: %s\n", diff.Key, formatValStylish(diff.NewVal, level+1))
		}
	}
	b.WriteString(strings.Repeat("  ", max(2*level-1, 0)) + "}")

	return b.String()
}

func formatValStylish(val any, level int) string {
	if m, ok := val.(map[string]any); ok {
		var b strings.Builder
		intent := strings.Repeat("  ", level*2+1)
		b.WriteString("{\n")

		keys := slices.Sorted(maps.Keys(m))

		for _, k := range keys {
			b.WriteString(intent)
			fmt.Fprintf(&b, "%s: %s\n", k, formatValStylish(m[k], level+1))
		}
		b.WriteString(strings.Repeat("  ", max(2*level-1, 0)) + "}")
		return b.String()
	}
	if val == nil {
		return "null"
	}
	return fmt.Sprint(val)
}
