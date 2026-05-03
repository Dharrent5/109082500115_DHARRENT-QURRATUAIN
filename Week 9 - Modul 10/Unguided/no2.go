package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("Masukkan jumlah ikan dan kapasitas wadah: ")
	fmt.Scan(&x, &y)

	ikan := make([]float64, x)
	for i := 0; i < x; i++ {
		fmt.Scan(&ikan[i])
	}

	var wadah []float646
	for i := 0; i < x; i += y {
		sum := 0.0
		for j := i; j < i+y && j < x; j++ {
			sum += ikan[j]
		}
		wadah = append(wadah, sum)
	}

	fmt.Println("Total berat tiap wadah:", wadah)

	total := 0.0
	for _, w := range wadah {
		total += w
	}
	fmt.Printf("Rata-rata berat wadah: %.2f\n", total/float64(len(wadah)))
}
