// Tipe Data Complex
// Tipe Data 	|	Deskripsi
// complex64	|	Bilangan kompleks dengan float32 bagian real dan imajiner
// complex128	|	Bilangan kompleks dengan float64 bagian real dan imajiner

package main

import (
	"fmt"
)

func main() {
	// Contoh penggunaan tipe data complex
	var bilangan1 complex64 = 1 + 2i
	var bilangan2 complex128 = 3.5 + 4.5i

	fmt.Println("Bilangan kompleks (complex64):", bilangan1)
	fmt.Println("Bilangan kompleks (complex128):", bilangan2)
}
