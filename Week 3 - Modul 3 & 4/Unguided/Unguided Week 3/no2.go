package main

import (
	"fmt"
)

func f(x int) int {
	return x * x
}

func g(x int) int {
	return x - 2
}

func h(x int) int {
	return x + 1
}

func main() {
	var a, b, c int

	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)
	fmt.Print("Masukkan nilai c: ")
	fmt.Scan(&c)

	result1 := f(g(h(a))) // (f ∘ g ∘ h)(a)
	result2 := g(h(f(b))) // (g ∘ h ∘ f)(b)
	result3 := h(f(g(c))) // (h ∘ f ∘ g)(c)

	fmt.Printf("(f ∘ g ∘ h)(a) = %d\n", result1)
	fmt.Printf("(g ∘ h ∘ f)(b) = %d\n", result2)
	fmt.Printf("(h ∘ f ∘ g)(c) = %d\n", result3)
}
