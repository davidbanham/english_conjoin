package english_conjoin

import (
	"fmt"
	"strings"
)

func Conjoin(items []string, joiner string) string {
	if len(items) == 0 {
		return ""
	}
	if len(items) == 1 {
		return items[0]
	}
	if len(items) == 2 { // "a and b" not "a, and b"
		return items[0] + " " + joiner + " " + items[1]
	}

	last, items := items[len(items)-1], items[:len(items)-1]
	secondLast, items := items[len(items)-1], items[:len(items)-1]

	items = append(items, secondLast+" "+joiner+" "+last)

	sep := ", "
	return strings.Join(items, sep)
}

func ConjoinAnd(items []string) string {
	return Conjoin(items, "and")
}

func ConjoinOr(items []string) string {
	return Conjoin(items, "or")
}

func DeDuplicate(items []string) []string {
	hits := map[string]int{}
	for _, line := range items {
		hits[line]++
	}
	output := []string{}
	fed := map[string]bool{}
	for _, line := range items {
		if fed[line] {
			continue
		}
		count := hits[line]
		newline := line
		if count > 1 {
			newline = fmt.Sprintf("%s (x%d)", line, hits[line])
		}
		output = append(output, newline)
		fed[line] = true
	}
	return output
}
