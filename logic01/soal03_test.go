package main

import (
	"fmt"
	"testing"
)

/*
0	1	2	3	4	5	6	7	8	9
1	99	2	99	3	99	4	99	5	99
*/

func TestSoal03(t *testing.T) {
	n := 10
	a := 1
	b := 99

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Print(a, "\t")
			a++
		} else {
			fmt.Print(b, "\t")
		}

	}
}
