package english_conjoin

import (
	"fmt"
	"strings"
)

func Conjoin(input []string, joiner string) string {
	if len(input) == 0 {
		return ""
	}
	if len(input) == 1 {
		return input[0]
	}
	if len(input) == 2 { // "a and b" not "a, and b"
		return input[0] + " " + joiner + " " + input[1]
	}

	items := make([]string, len(input))
	copy(items, input)

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

func DeDuplicate(input []string) []string {
	items := make([]string, len(input))
	copy(items, input)

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
