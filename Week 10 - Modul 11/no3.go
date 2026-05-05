package main

import "fmt"

const nProv int = 10

func InputData(prov []string, pop []int, tumbuh []float64) {
	for i := 0; i < nProv; i++ {
		fmt.Printf("Masukkan data ke-%d (Nama Populasi Pertumbuhan): ", i+1)
		fmt.Scan(&prov[i], &pop[i], &tumbuh[i])
	}
}

func ProvinsiTercepat(tumbuh []float64) int {
	idx := 0
	for i := 1; i < nProv; i++ {
		if tumbuh[i] > tumbuh[idx] {
			idx = i
		}
	}
	return idx
}

func IndeksProvinsi(prov []string, nama string) int {
	for i := 0; i < nProv; i++ {
		if prov[i] == nama {
			return i
		}
	}
	return -1
}

func Prediksi(prov []string, pop []int, tumbuh []float64) {
	fmt.Println("Data Provinsi dengan pertumbuhan di atas 2%:")
	for i := 0; i < nProv; i++ {
		if tumbuh[i] > 0.02 {
			prediksi := int((1 + tumbuh[i]) * float64(pop[i]))
			fmt.Println(prov[i], prediksi)
		}
	}
}

func main() {
	prov := make([]string, nProv)
	pop := make([]int, nProv)
	tumbuh := make([]float64, nProv)

	fmt.Println("Masukkan Nama Provinsi, Populasi, Pertumbuhan:")
	InputData(prov, pop, tumbuh)

	var cari string
	fmt.Print("Masukkan nama provinsi yang ingin dicari: ")
	fmt.Scan(&cari)

	idxCepat := ProvinsiTercepat(tumbuh)
	fmt.Println("Provinsi dengan pertumbuhan tercepat:", prov[idxCepat])

	idxCari := IndeksProvinsi(prov, cari)
	if idxCari != -1 {
		fmt.Println("Data provinsi yang dicari:", prov[idxCari])
	} else {
		fmt.Println("Provinsi tidak ditemukan")
	}

	Prediksi(prov, pop, tumbuh)
}
