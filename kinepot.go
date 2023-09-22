package main

import (
	"fmt"
)

func main() {
	fmt.Print("Masukkan massa benda (kg): ")
	var massa float64
	fmt.Scanln(&massa)

	fmt.Print("Masukkan ketinggian (m): ")
	var ketinggian float64
	fmt.Scanln(&ketinggian)
	
	fmt.Print("Masukkan kecepatan (mps): ")
	var kecepatan float64
	fmt.Scanln(&kecepatan)

	g := 9.81 

	energiPotensial := massa * g * ketinggian

	energiKinetik := massa * kecepatan

	fmt.Printf("Energi Potensial: %.2f joule\n", energiPotensial)
	fmt.Printf("Energi Kinetik: %.2f joule\n", energiKinetik)
}
