package english_conjoin

import "strings"

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
