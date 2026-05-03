package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan jumlah kelinci: ")
	fmt.Scan(&n)

	berat := make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan berat kelinci ke-%d: ", i+1)
		fmt.Scan(&berat[i])
	}

	min, max := berat[0], berat[0]
	for _, b := range berat {
		if b < min {
			min = b
		}
		if b > max {
			max = b
		}
	}

	fmt.Printf("Berat terkecil: %.2f\n", min)
	fmt.Printf("Berat terbesar: %.2f\n", max)
}
