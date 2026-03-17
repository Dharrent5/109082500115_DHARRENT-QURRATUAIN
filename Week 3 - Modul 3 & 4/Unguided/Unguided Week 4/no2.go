package main

import (
	"fmt"
	"math"
)

func hitungPersegi(sisi int) {
	luas := sisi * sisi
	keliling := 4 * sisi
	fmt.Printf("Persegi dengan sisi %d\n", sisi)
	fmt.Printf("Luas: %d\n", luas)
	fmt.Printf("Keliling: %d\n", keliling)
}

func hitungPersegiPanjang(panjang, lebar int) {
	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)
	fmt.Printf("Persegi Panjang dengan panjang %d dan lebar %d\n", panjang, lebar)
	fmt.Printf("Luas: %d\n", luas)
	fmt.Printf("Keliling: %d\n", keliling)
}

func hitungLingkaran(jariJari float64) {
	luas := math.Pi * jariJari * jariJari
	keliling := 2 * math.Pi * jariJari
	fmt.Printf("Lingkaran dengan jari-jari %.2f\n", jariJari)
	fmt.Printf("Luas: %.2f\n", luas)
	fmt.Printf("Keliling: %.2f\n", keliling)
}

func main() {
	var pilihan int
	fmt.Println("--- PROGRAM BANGUN DATAR ---")
	fmt.Println("1. Hitung luas & keliling persegi")
	fmt.Println("2. Hitung luas & keliling persegi panjang")
	fmt.Println("3. Hitung luas & keliling lingkaran")
	fmt.Print("Pilihan : ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		hitungPersegi(1057)
	case 2:
		hitungPersegiPanjang(106, 88)
	case 3:
		hitungLingkaran(13.87)
	default:
		fmt.Println("Pilihan tidak valid")
	}
}
