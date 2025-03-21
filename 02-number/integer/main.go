// Tipe Data Integer
// Tipe Data 	|	Nilai Minimum			|	Nilai Maksimum		|
// int8			|	-128					|	127					|
// int16		|	-32768					|	32767				|
// int32		|	-2147483648				|	2147483647			|
// int64		|	-9223372036854775808	|	9223372036854775807	|

// ========================================================================

// Tipe Data Unsigned Integer
// Tipe Data 	|	Nilai Minimum			|	Nilai Maksimum			|
// uint8		|	0						|	255						|
// uint16		|	0						|	65535					|
// uint32		|	0						|	4294967295				|
// uint64		|	0						|	18446744073709551615	|

package main

import "fmt"

func main() {
	// Contoh deklarasi variabel tanpa tipe data
	var x = 42
	var y = 3.14
	var z = "Hello, Golang!"

	fmt.Println("Contoh tanpa deklarasi tipe data:")
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)

	// Contoh penggunaan tipe data integer
	var a int8 = 127
	var b int16 = 32767
	var c int32 = 2147483647
	var d int64 = 9223372036854775807

	fmt.Println("\nContoh Integer:")
	fmt.Println("int8:", a)
	fmt.Println("int16:", b)
	fmt.Println("int32:", c)
	fmt.Println("int64:", d)

	// Contoh penggunaan tipe data unsigned integer
	var ua uint8 = 255
	var ub uint16 = 65535
	var uc uint32 = 4294967295
	var ud uint64 = 18446744073709551615

	fmt.Println("\nContoh Unsigned Integer:")
	fmt.Println("uint8:", ua)
	fmt.Println("uint16:", ub)
	fmt.Println("uint32:", uc)
	fmt.Println("uint64:", ud)
}
