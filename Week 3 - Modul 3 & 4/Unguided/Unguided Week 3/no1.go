package main

import (
	"fmt"
)

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func permutasi(n, r int) int {
	return factorial(n) / factorial(n-r)
}

func kombinasi(n, r int) int {
	return factorial(n) / (factorial(r) * factorial(n-r))
}

func main() {
	var a, b, c, d int

	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)
	fmt.Print("Masukkan nilai c: ")
	fmt.Scan(&c)
	fmt.Print("Masukkan nilai d: ")
	fmt.Scan(&d)

	fmt.Printf("Permutasi a terhadap c: %d\n", permutasi(a, c))
	fmt.Printf("Kombinasi a terhadap c: %d\n", kombinasi(a, c))
	fmt.Printf("Permutasi b terhadap d: %d\n", permutasi(b, d))
	fmt.Printf("Kombinasi b terhadap d: %d\n", kombinasi(b, d))
}
