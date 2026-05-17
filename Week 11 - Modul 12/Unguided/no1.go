package main

import "fmt"

func main() {
	var suaraMasuk []int
	var x int
	fmt.Println("Masukkan data suara (akhiri dengan 0):")
	for {
		fmt.Scan(&x)
		if x == 0 {
			break
		}
		suaraMasuk = append(suaraMasuk, x)
	}

	total := len(suaraMasuk)
	sah := 0
	hasil := make([]int, 21) // indeks 1..20
	for _, v := range suaraMasuk {
		if v >= 1 && v <= 20 {
			hasil[v]++
			sah++
		}
	}

	fmt.Printf("Suara masuk: %d\n", total)
	fmt.Printf("Suara sah: %d\n", sah)
	for i := 1; i <= 20; i++ {
		if hasil[i] > 0 {
			fmt.Printf("%d: %d\n", i, hasil[i])
		}
	}
}
