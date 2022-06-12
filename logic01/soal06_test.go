package main

import (
	"fmt"
	"math"
	"testing"
)

func TestSoalNo6(t *testing.T) {
	n := 15
	nt := n / 2
	x := 3
	for i := 1; i <= n; i++ {
		if i%4 == 0 {
			z := math.Pow(float64(i), 2)
			fmt.Print(z, "\t")
		} else {
			fmt.Print(x, "\t")
		}
		if i <= nt {
			x += 3
		} else {
			x -= 3
		}
	}
}
