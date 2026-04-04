package main
import "fmt"
5
func faktor(n, i int) {
    if i > n {
        return
    }
    if n%i == 0 {
        fmt.Print(i, " ")
    }
    faktor(n, i+1)
}

func main() {
    var n int
    fmt.Print("Masukkan bilangan N: ")
    fmt.Scan(&n)

    fmt.Println("Faktor dari", n, ":")
    faktor(n, 1)
}
