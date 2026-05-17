package main

import "fmt"

func main() {
	var suara []int
	var x int
	for {
		fmt.Scan(&x)
		if x == 0 {
			break
		}
		suara = append(suara, x)
	}

	total := len(suara)
	sah := 0
	hasil := make([]int, 21)
	for _, v := range suara {
		if v >= 1 && v <= 20 {
			hasil[v]++
			sah++
		}
	}

	fmt.Printf("Suara masuk: %d\n", total)
	fmt.Printf("Suara sah: %d\n", sah)

	ketua, wakil := 0, 0
	for i := 1; i <= 20; i++ {
		if hasil[i] > hasil[ketua] || (hasil[i] == hasil[ketua] && i < ketua) {
			wakil = ketua
			ketua = i
		} else if hasil[i] > hasil[wakil] || (hasil[i] == hasil[wakil] && i < wakil && i != ketua) {
			wakil = i
		}
	}

	fmt.Printf("Ketua RT: %d\n", ketua)
	fmt.Printf("Wakil ketua: %d\n", wakil)
}
