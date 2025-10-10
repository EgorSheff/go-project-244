package formatters

import (
	"code/getdiff"
	"fmt"
	"strings"
)

func FormatDiffsPlain(diffs []getdiff.Diff) string {
	return strings.TrimSpace(formatPlain(diffs, ""))
}

func formatPlain(diffs []getdiff.Diff, root string) string {
	var b strings.Builder
	for _, diff := range diffs {
		k := diff.Key
		if root != "" {
			k = root + "." + diff.Key
		}

		if len(diff.NestDiffs) > 0 {
			b.WriteString(formatPlain(diff.NestDiffs, k))
			continue
		}

		if diff.Equal() {
			continue
		}

		if diff.Del && diff.Add {
			fmt.Fprintf(&b, "Property '%s' was updated. From %s to %s\n", k, formatValPlain(diff.OldVal), formatValPlain(diff.NewVal))
			continue
		}

		if diff.Del {
			fmt.Fprintf(&b, "Property '%s' was removed\n", k)
		}

		if diff.Add {
			fmt.Fprintf(&b, "Property '%s' was added with value: %s\n", k, formatValPlain(diff.NewVal))
		}
	}
	return b.String()
}

func formatValPlain(val any) string {
	if _, ok := val.(map[string]any); ok {
		return "[complex value]"
	}

	switch val.(type) {
	case string:
		return fmt.Sprintf("'%s'", val)
	case nil:
		return "null"
	}
	return fmt.Sprint(val)
}
