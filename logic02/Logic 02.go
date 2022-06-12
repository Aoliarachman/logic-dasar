package main

import (
	"fmt"
)

func Logic02soal01(n int) {

	for i := 0; i < n; i++ { //loop baris ke pinggir
		a := 3
		for j := 0; j < n; j++ { //loop kolom ke bawah
			fmt.Print(a, "\t")
			a = 3
		}
		fmt.Println("\n")
	}

}

func Logic02soal02(n int) {

	for i := 0; i < n; i++ { //loop baris ke pinggir
		a := 3
		for j := 0; j < n; j++ { //loop kolom ke bawah
			fmt.Print(a, "\t")
			a += 3
		}
		fmt.Println("\n")
	}
}

func Logic02soal03(n int) {
	for i := 0; i < n; i++ {
		a := 3 * n
		for j := 0; j < n; j++ {
			fmt.Print(a, "\t")
			a -= 3
		}
		fmt.Println("\n")

	}
}

func Logic02soal04(n int) {
	a := 3 * n // nilai variable
	// i sebagai loop baris
	for i := 0; i < n; i++ {
		// j sebagai loop kolom
		for j := 0; j < n; j++ {
			//if i == j
			fmt.Print(a, "\t")
		}
		//kolom selesai dan selanjutnya ke baris
		fmt.Println("\n")
		a -= 3 // tambahan nilai untuk nilai a
	}
}

func Logic02soal05(n int) {
	//nilai tengah
	nt := n / 2
	//nilai yang akan bertambah
	a := 3

	//looping baris
	for i := 0; i < n; i++ {
		//looping kolom
		for j := 0; j < n; j++ {
			fmt.Print(a, "\t")
		} // kolom selesai
		// ke baris selanjutnya :
		fmt.Println("\n")
		//update value variable a
		if i < nt {
			a += 3
		} else {
			a -= 3
		}

	}
}

func Logic02soal06(n int) {
	nt := n / 2

	//looping baris
	for i := 0; i < n; i++ {
		//set value variable
		a := 3
		// looping kolom
		for j := 0; j < n; j++ {
			if j < nt {
				fmt.Print(a, "\t")
				//update value variable a
				a += 3
			} else {
				fmt.Print(a, "\t")
				//update value variable a
				a -= 3
			}
		}
		//pindah baris
		fmt.Println("\n")
	}
}

func Logic02soal07(n int) {
	a := 3 // nilai variable
	// i sebagai loop baris
	for i := 0; i < n; i++ {
		// j sebagai loop kolom
		for j := 0; j < n; j++ {
			if i < j {
				fmt.Print(" ", "\t")
			} else {
				fmt.Print(a, "\t")
			}
			// kolom selesai dan melanjutkan ke baris
			fmt.Println("\n")
			a += 3 // tambahan nilai utk nilai a

		}
	}
}

func Logic02soal08(n int) {
	a := 3 // nilai variable
	// i sebagai loop baris
	for i := 0; i < n; i++ {
		// j sebagai loop kolom
		for j := 0; j < n; j++ {
			if i > j {
				fmt.Print(" ", "\t")
			} else {
				fmt.Print(a, "\t")
			}
		}
		// kolom selesai dan melanjutkan ke baris
		fmt.Println("\n")
		a += 3 // tambahan nilai utk nilai a
	}
}

func Logic02soal09(n int) {
	// nilai variable
	// i sebagai loop baris
	for i := 0; i < n; i++ {
		a := 0 // nilai 27
		// j sebagai loop kolom
		for j := 0; j < n; j++ {
			a += 3
			if i+j == n-1 {
				fmt.Print(a, "\t")
			} else if i+j >= n-1 {
				fmt.Print(a, "\t")
			} else {
				fmt.Print(" ", "\t")
			}
		}
		// kolom selesai dan melanjutkan ke baris
		fmt.Println("\n")
		//a += 3 // tambahan nilai utk nilai a
	}
}

func Logic02soal10(n int) {
	for i := 0; i < n; i++ {
		a := 0 // nilai 27
		// j sebagai loop kolom
		for j := 0; j < n; j++ {
			a += 3
			if i+j == n-1 {
				fmt.Print(a, "\t")
			} else if i+j <= n-1 {
				fmt.Print(a, "\t")
			} else {
				fmt.Print(" ", "\t")
			}
		}
		// kolom selesai dan melanjutkan ke baris
		fmt.Println("\n")
		//a += 3 // tambahan nilai utk nilai a
	}
}
