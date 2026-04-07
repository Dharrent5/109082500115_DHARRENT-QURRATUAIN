package main

import "fmt"

func tanggunganHari(jumlahHari int, tujuan string) int {
	var maxHari int
	if tujuan == "domestik" {
		maxHari = 3
	} else if tujuan == "mancanegara" {
		maxHari = 8
	}

	if jumlahHari > maxHari {
		return maxHari
	}
	return jumlahHari
}

func biayaPerHari(jumlahMhs int) int {
	biayaSatuMhs := (2 * 35000) + 250000 + 300000
	return biayaSatuMhs * jumlahMhs
}

func perhitunganBiaya(jumlahMhs, lamaPerjalanan int, tujuan string, totalBiaya *float64) {
	hariDitanggung := tanggunganHari(lamaPerjalanan, tujuan)

	biayaDomestikHarian := biayaPerHari(jumlahMhs)

	if tujuan == "mancanegara" {
		*totalBiaya = float64(hariDitanggung) * float64(biayaDomestikHarian) * 1.5
	} else {
		*totalBiaya = float64(hariDitanggung) * float64(biayaDomestikHarian)
	}
}

func main() {
	var jumlah, lama int
	var tujuan string
	var biaya float64

	fmt.Print("masukkan jumlah mahasiswa : ")
	fmt.Scan(&jumlah)
	fmt.Print("masukkan lama hari study tour : ")
	fmt.Scan(&lama)
	fmt.Print("masukkan tujuan study tour (domestik/mancanegara) : ")
	fmt.Scan(&tujuan)

	perhitunganBiaya(jumlah, lama, tujuan, &biaya)

	fmt.Printf("\nBiaya perjalanan yang harus dikeluarkan Tel-U : Rp. %g\n", biaya)
}
