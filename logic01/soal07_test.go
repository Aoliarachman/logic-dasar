package main

import (
	"fmt"
	"testing"
)

func TestSoal07(t *testing.T) {
	n := 10
	a := 4
	b := 3

	for i := 1; i <= n; i++ {
		if i%2 == 1 {
			fmt.Print(a, "\t")

		} else {
			fmt.Print(b, "\t")
			b += 3
		}

	}
}
