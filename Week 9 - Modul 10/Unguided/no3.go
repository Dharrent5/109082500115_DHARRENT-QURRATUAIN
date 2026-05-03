package main

import "fmt"

func hitungMinMax(arr []float64) (float64, float64) {
	min, max := arr[0], arr[0]
	for _, v := range arr {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max
}

func rerata(arr []float64) float64 {
	sum := 0.0
	for _, v := range arr {
		sum += v
	}
	return sum / float64(len(arr))
}

func main() {
	var n int
	fmt.Print("Masukkan banyak data balita: ")
	fmt.Scan(&n)

	arr := make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan berat balita ke-%d: ", i+1)
		fmt.Scan(&arr[i])
	}

	min, max := hitungMinMax(arr)
	fmt.Printf("Berat balita minimum: %.2f kg\n", min)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", max)
	fmt.Printf("Rata-rata berat balita: %.2f kg\n", rerata(arr))
}
