package main

import (
	"fmt"
)

func main() {
	correctOrder := []string{"merah", "kuning", "hijau", "ungu"}

	allSuccess := true

	for i := 1; i <= 5; i++ {
		fmt.Printf("Masukkan warna untuk percobaan %d (4 warna dipisahkan spasi): ", i)

		var w1, w2, w3, w4 string
		fmt.Scan(&w1, &w2, &w3, &w4)

		if w1 != correctOrder[0] || w2 != correctOrder[1] || w3 != correctOrder[2] || w4 != correctOrder[3] {
			allSuccess = false
		}
	}

	fmt.Printf("BERHASIL : %t\n", allSuccess)
}
