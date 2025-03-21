// Tipe Data Floating Point
// Tipe Data 	|	Deskripsi
// float32		|	Bilangan real dengan presisi hingga 7 digit desimal
// float64		|	Bilangan real dengan presisi hingga 15 digit desimal

package main

import "fmt"

func main() {
	// Contoh penggunaan tipe data floating point
	var panjang float32 = 12.34
	var lebar float32 = 56.78
	var luas float32 = panjang * lebar
	fmt.Println("Luas persegi panjang:", luas)

	var radius float64 = 5.67
	var pi float64 = 3.14159
	var area float64 = pi * (radius * radius)
	fmt.Println("Luas lingkaran:", area)
}
