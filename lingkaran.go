package main

import (
    "fmt"
    "math"
)

func main() {

    fmt.Print("Masukkan jari-jari lingkaran: ")
    var jariJari float64
    fmt.Scanln(&jariJari)

    luas := math.Pi * jariJari * jariJari

    fmt.Printf("Luas lingkaran  adalah %.2f\n", luas)
}
