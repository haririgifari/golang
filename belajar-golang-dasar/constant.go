// Constant
// 1. Constant adalah variable yang nilainya tidak bisa diubah lagi setelah pertama kali diberi nilai
// 2. Cara pembuatan constat mirip dengan variable yang membedakan hanya kata kunci yang digunakan adalah const, bukan var
// 3. Saat pembuatan constant, kita wajib langsung menginisialisasikan datanya

package main

import "fmt"

func main() {
	const firstName string = "Hariri"
	const lastName = "Gifari"
	const value = 1000

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)
}

// Deklarasi multiple constant
// Sama seperti variable, di glang juga kita bisa membuat constant secara sekaligus banyak

package main

import "fmt"

func main() {
	const (
		firstName string = "Hariri"
		lastName         = "Gifari"
		Value            = 1000
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(Value)
}
