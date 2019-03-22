package english_conjoin

import (
	"strings"
	"testing"
)

type fixture struct {
	input  []string
	result string
}

func TestConjoiner(t *testing.T) {
	data := []fixture{
		{
			input:  []string{},
			result: "",
		},
		{
			input:  []string{""},
			result: "",
		},
		{
			input:  []string{"Alice"},
			result: "Alice",
		},
		{
			input:  []string{"Alice", "Bob"},
			result: "Alice and Bob",
		},
		{
			input:  []string{"Alice", "Bob", "Charlie"},
			result: "Alice, Bob and Charlie",
		},
		{
			input:  []string{"Alice", "Bob", "Charlie", "Denise"},
			result: "Alice, Bob, Charlie and Denise",
		},
	}
	for _, fixture := range data {
		t.Run(strings.Join(fixture.input, " ")+" -> "+fixture.result, func(t *testing.T) {
			result := Conjoin(fixture.input, "and")
			if result != fixture.result {
				t.Errorf("got %s, want %s", result, fixture.result)
			}
		})
	}
}
