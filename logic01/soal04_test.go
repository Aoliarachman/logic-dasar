package main

import (
	"fmt"
	"testing"
)

func TestSoal04(t *testing.T) {
	n := 10
	a := 1
	b := 177

	for i := 0; i < n; i++ {
		if i%2 == 1 {
			fmt.Print(a, "\t")
			a++
		} else {
			fmt.Print(b, "\t")
		}

	}
}
