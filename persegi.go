package main

import "fmt"

func main() {
    fmt.Print("Masukkan panjang sisi : ")
    var sisi float64
    fmt.Scanln(&sisi)
    luas := sisi * sisi
    fmt.Printf("Luas persegi dengan sisi %.2f adalah %.2f\n", sisi, luas)
}
