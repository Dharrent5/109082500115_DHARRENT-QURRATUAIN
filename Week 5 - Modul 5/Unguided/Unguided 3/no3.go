package main
import "fmt"

func ganjil(n, i int) {
    if i > n {
        return
    }
    fmt.Print(i, " ")
    ganjil(n, i+2)
}

func main() {
    var n int
    fmt.Print("Masukkan bilangan N: ")
    fmt.Scan(&n)

    fmt.Println("Bilangan ganjil dari 1 sampai", n, ":")
    ganjil(n, 1)
}
