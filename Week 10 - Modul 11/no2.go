package main

import "fmt"

type mahasiswa struct {
	NIM   string
	nama  string
	nilai int
}

func nilaiPertama(arr []mahasiswa, n int, nim string) int {
	for i := 0; i < n; i++ {
		if arr[i].NIM == nim {
			return arr[i].nilai
		}
	}
	return -1
}

func nilaiTerbesar(arr []mahasiswa, n int, nim string) int {
	max := -1
	for i := 0; i < n; i++ {
		if arr[i].NIM == nim {
			if arr[i].nilai > max {
				max = arr[i].nilai
			}
		}
	}
	return max
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah data: ")
	fmt.Scan(&n)

	data := make([]mahasiswa, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan data ke-%d (NIM Nama Nilai): ", i+1)
		fmt.Scan(&data[i].NIM, &data[i].nama, &data[i].nilai)
	}

	var cariNIM string
	fmt.Print("Masukkan NIM mahasiswa yang ingin dicari: ")
	fmt.Scan(&cariNIM)

	pertama := nilaiPertama(data, n, cariNIM)
	terbesar := nilaiTerbesar(data, n, cariNIM)

	if pertama == -1 {
		fmt.Println("Data mahasiswa dengan NIM", cariNIM, "tidak ditemukan.")
	} else {
		fmt.Printf("Nilai pertama dari NIM %s adalah %d\n", cariNIM, pertama)
		fmt.Printf("Nilai terbesar dari NIM %s adalah %d\n", cariNIM, terbesar)
	}
}
