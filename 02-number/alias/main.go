// Alias Tipe Data
// Alias adalah nama lain untuk tipe data yang sudah ada, digunakan untuk memberikan nama yang lebih deskriptif atau sesuai konteks.

package main

import "fmt"

// Mendefinisikan alias untuk tipe data
type Kilometer float64
type Liter float64
type Second int

func main() {
	// Menggunakan alias tipe data
	var jarak Kilometer = 15.5
	var bahanBakar Liter = 5.2
	var waktu Second = 3600

	// Menggunakan alias tipe data tambahan
	var Byte byte = 255 // uint8
	var Rune rune = 100 // int32
	var Int int = -200  // Minimal int32
	var Uint uint = 300 // Minimal uint32

	fmt.Println("Data byte:", Byte)
	fmt.Println("Data rune:", Rune)
	fmt.Println("Data int:", Int)
	fmt.Println("Data uint:", Uint)

	fmt.Println("Jarak tempuh:", jarak, "kilometer")
	fmt.Println("Konsumsi bahan bakar:", bahanBakar, "liter")
	fmt.Println("Waktu tempuh:", waktu, "detik")
}
