package main

import (
	"fmt"
	"testing"
)

//setiap kelipat 3 fizz
//setiap kelipatan 5 buzz
//kelipatan 3&5 fizz buzz

func TestSoal05(t *testing.T) {
	n := 15
	//i kurang dr n artinya perulangan dihentikan ketikan nilai i kurang dr n
	//for i := 1; (angka 1
	for i := 1; i <= n; i++ { // = untuk memunculkan semua angka
		if i%3 == 0 && i%5 == 0 {
			fmt.Print("fizz buzz", "\t")
		} else if i%3 == 0 { // i%3 == 0 artiny jika i habis dibagi 3 (true) maka eksekusi line 20. jika false masuk ke else
			fmt.Print("fizz", "\t")
		} else if i%5 == 0 {
			fmt.Print("buzz", "\t")
		} else {
			fmt.Print(i, "\t")
		}

	}

}
