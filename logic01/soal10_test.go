package main

import (
	"fmt"
	"testing"
)

func TestSoal10(t *testing.T) {
	n := 12
	a := 1
	b := 2
	c := 3

	for i := 1; i <= n; i++ {
		if i%4 == 1 {
			fmt.Print(c, "\t")
			c += 3
		} else if i%4 == 2 {
			fmt.Print(b, "\t")
			b += 2
		} else if i%4 == 3 {
			fmt.Print(a, "\t")
			a += 1
		} else {
			fmt.Print(999, "\t")
		}

	}
}
