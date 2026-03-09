package main

import (
	"fmt"
)

func isLeapYear(year int) bool {
	if year%400 == 0 {
		return true
	} else if year%4 == 0 && year%100 != 0 {
		return true
	}
	return false
}

func main() {
	var year int
	fmt.Print("Masukkan tahun: ")
	fmt.Scan(&year)

	if isLeapYear(year) {
		fmt.Printf("Tahun %d adalah tahun kabisat.\n", year)
	} else {
		fmt.Printf("Tahun %d bukan tahun kabisat.\n", year)
	}
}
