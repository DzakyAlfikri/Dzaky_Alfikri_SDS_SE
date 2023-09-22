package main

import (
	"fmt"
)

func main() {
	var suhu float64
	var pilihan int

	fmt.Println("Program Konversi Satuan Suhu")
	fmt.Println("1. Celcius ke Fahrenheit")
	fmt.Println("2. Celcius ke Kelvin")
	fmt.Println("3. Fahrenheit ke Celcius")
	fmt.Println("4. Fahrenheit ke Kelvin")
	fmt.Println("5. Kelvin ke Celcius")
	fmt.Println("6. Kelvin ke Fahrenheit")
	fmt.Print("Pilih opsi (1/2/3/4/5/6): ")
	fmt.Scanln(&pilihan)

	fmt.Print("Masukkan suhu: ")
	fmt.Scanln(&suhu)

	if pilihan == 1 {
		celciusToFahrenheit(suhu)
	} else if pilihan == 2 {
		celciusToKelvin(suhu)
	} else if pilihan == 3 {
		fahrenheitToCelcius(suhu)
	} else if pilihan == 4 {
		fahrenheitToKelvin(suhu)
	} else if pilihan == 5 {
		kelvinToCelcius(suhu)
	} else if pilihan == 6 {
		kelvinToFahrenheit(suhu)
	} else {
		fmt.Println("Pilihan tidak valid")
	}
}

func celciusToFahrenheit(c float64) {
	f := (c * 9/5) + 32
	fmt.Printf("%.2f Celcius = %.2f Fahrenheit\n", c, f)
}

func celciusToKelvin(c float64) {
	k := c + 273.15
	fmt.Printf("%.2f Celcius = %.2f Kelvin\n", c, k)
}

func fahrenheitToCelcius(f float64) {
	c := (f - 32) * 5/9
	fmt.Printf("%.2f Fahrenheit = %.2f Celcius\n", f, c)
}

func fahrenheitToKelvin(f float64) {
	k := (f-32)*5/9 + 273.15
	fmt.Printf("%.2f Fahrenheit = %.2f Kelvin\n", f, k)
}

func kelvinToCelcius(k float64) {
	c := k - 273.15
	fmt.Printf("%.2f Kelvin = %.2f Celcius\n", k, c)
}

func kelvinToFahrenheit(k float64) {
	f := (k-273.15)*9/5 + 32
	fmt.Printf("%.2f Kelvin = %.2f Fahrenheit\n", k, f)
}
