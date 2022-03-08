package english_conjoin

import (
	"strings"
	"testing"
)

type fixture struct {
	input  []string
	result string
}

type dupeFixture struct {
	input  []string
	result []string
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

func TestDeDuplicator(t *testing.T) {
	data := []dupeFixture{
		{
			input:  []string{},
			result: []string{},
		},
		{
			input:  []string{""},
			result: []string{""},
		},
		{
			input:  []string{"Alice"},
			result: []string{"Alice"},
		},
		{
			input:  []string{"Alice", "Bob", "Bob"},
			result: []string{"Alice", "Bob (x2)"},
		},
		{
			input:  []string{"Alice", "Bob", "Charlie"},
			result: []string{"Alice", "Bob", "Charlie"},
		},
		{
			input:  []string{"Milk", "Flour", "Eggs", "Rice", "Flour", "Rice", "Rice"},
			result: []string{"Milk", "Flour (x2)", "Eggs", "Rice (x3)"},
		},
	}
	for _, fixture := range data {
		t.Run(strings.Join(fixture.input, " ")+" -> "+strings.Join(fixture.result, " "), func(t *testing.T) {
			result := DeDuplicate(fixture.input)
			if len(result) == len(fixture.result) {
				for i, resPart := range result {
					if fixture.result[i] != resPart {
						t.Errorf("got %s, want %s", result, fixture.result)
					}
				}
			} else {
				t.Errorf("got %s, want %s", result, fixture.result)
			}
		})
	}
}
