package main

import "fmt"

func main() {
    fmt.Print("Masukkan panjang alas segitiga: ")
    var alas float64
    fmt.Scanln(&alas)

    fmt.Print("Masukkan tinggi segitiga: ")
    var tinggi float64
    fmt.Scanln(&tinggi)

    luas := 0.5 * alas * tinggi

    fmt.Printf("Luas segitiga adalah %.2f\n", luas)
}
