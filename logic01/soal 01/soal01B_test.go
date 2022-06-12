package soal_01

import (
	"fmt"
	"testing"
)

func TestSoal01B(t *testing.T) {

	//looping artinya pengulangan
	n := 10
	// berarti + =1
	angka := 1
	// for artinya looping , balik ke program sampai ke-n kali
	// for i := 0 artinya looping di mulai dari 0
	for i := 0; i < n; i++ {
		// if 1%2 == 0 artinya jika i di bagi 2 hasilnya 0 makan true
		// if true masuk ke kurung kurawal pertama if
		// if true fmt.Println(angka, "\t") angka ++
		// if false yang di dalam kurang skip
		// if 1%2 == 0 habis di bagi 2 (genap) makan print angka
		// modulus 2 hasilnya 0 makan genap
		if i%2 == 0 {
			fmt.Print(angka, "\t")
			angka++
		} else {
			fmt.Print(3, "\t")
			//jk index genap variable "angka" +1
		}
	}
}
