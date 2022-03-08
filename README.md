# English Conjoin[![Build Status](https://travis-ci.org/davidbanham/english_conjoin.svg?branch=master)](https://travis-ci.org/davidbanham/english_conjoin) [![GoDoc](https://godoc.org/github.com/davidbanham/english_conjoin?status.svg)](https://godoc.org/github.com/davidbanham/english_conjoin)
A little Go util to print lists of strings in common English format

## Docs

The [Conjoin](https://godoc.org/github.com/davidbanham/english_conjoin#Conjoin) function takes a slice of strings and a final joiner.

## Usage

```go
package main

import (
	"fmt"
	"github.com/davidbanham/english_conjoin"
)

func main() {
	example := []string{"Alice", "Bob", "Charlie"}

	fmt.Println(english_conjoin.Conjoin(example, "and")) // Alice, Bob and Charlie
	fmt.Println(english_conjoin.Conjoin(example, "or"))  // Alice, Bob or Charlie

	dupeExample := []string{"Milk", "Flour", "Eggs", "Rice", "Flour", "Rice", "Rice"}

	fmt.Println(english_conjoin.DeDuplicate(dupeExample))  // Milk, Flour (x2), Eggs, Rice (x3)
}
```

There are also some helper funcs for common joiners: `ConjoinAnd` and `ConjoinOr`

There are more examples in the [tests](https://github.com/davidbanham/english_conjoin/blob/master/main_test.go).
