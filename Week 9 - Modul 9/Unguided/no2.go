package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Masukkan jumlah elemen: ")
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Println("Isi array:", arr)

	fmt.Print("Indeks ganjil: ")
	for i := 1; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	fmt.Print("Indeks genap: ")
	for i := 0; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	var x int
	fmt.Print("Masukkan kelipatan indeks x: ")
	fmt.Scan(&x)
	fmt.Print("Indeks kelipatan ", x, ": ")
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	var del int
	fmt.Print("Masukkan indeks yang dihapus: ")
	fmt.Scan(&del)
	arr = append(arr[:del], arr[del+1:]...)
	fmt.Println("Array setelah dihapus:", arr)

	sum := 0
	for _, v := range arr {
		sum += v
	}
	mean := float64(sum) / float64(len(arr))
	fmt.Println("Rata-rata:", mean)

	var sd float64
	for _, v := range arr {
		sd += math.Pow(float64(v)-mean, 2)
	}
	sd = math.Sqrt(sd / float64(len(arr)))
	fmt.Println("Standar deviasi:", sd)

	var cari int
	fmt.Print("Masukkan bilangan untuk frekuensi: ")
	fmt.Scan(&cari)
	freq := 0
	for _, v := range arr {
		if v == cari {
			freq++
		}
	}
	fmt.Println("Frekuensi:", freq)
}
