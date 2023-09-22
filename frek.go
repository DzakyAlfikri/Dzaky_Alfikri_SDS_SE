package main

import (
	"fmt"
)

func main() {
	fmt.Print("Masukkan periode getaran (detik): ")
	var periode float64
	fmt.Scanln(&periode)
	
	frekuensi := 1.0 / periode

	fmt.Printf("Frekuensi getaran adalah %.2f Hz\n", frekuensi)
}
