package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-1-i] = t[n-1-i], t[i]
	}
}

func palindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var n int
	fmt.Print("Masukkan kata: ")
	var s string
	fmt.Scan(&s)
	n = len(s)
	for i, r := range s {
		tab[i] = r
	}

	fmt.Print("Balikan: ")
	balikanArray(&tab, n)
	for i := 0; i < n; i++ {
		fmt.Print(string(tab[i]))
	}
	fmt.Println()

	if palindrom(tab, n) {
		fmt.Println("Palindrom")
	} else {
		fmt.Println("Bukan palindrom")
	}
}
