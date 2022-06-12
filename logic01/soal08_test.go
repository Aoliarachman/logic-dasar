package main

import (
	"fmt"
	"testing"
)

func TestSoal08(t *testing.T) {
	n := 10
	angkakelipatan2 := 2
	angkakelipatan5 := 5

	for i := 1; i <= n; i++ {
		if i%2 == 1 {
			fmt.Print(angkakelipatan2, "\t")
			angkakelipatan2 *= 2
		} else {
			fmt.Print(angkakelipatan5, "\t")
			angkakelipatan5 *= 5
		}

	}
}
