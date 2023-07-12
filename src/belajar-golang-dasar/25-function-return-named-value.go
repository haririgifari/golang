// Named Return Values
// 1. Biasanya saat kita memberi tahu bahwa sebuah function mengembalikan value, maka kita hanya mendeklarasika tipe data return value di function
// 2. Namun kita juga bisa membuat variable secara langsung di tipe data return function nya

package main

import "fmt"

// kalau ingin pakai bolean atau int tinggal belakangnya di ikuti
// contoh 22 int dll (yang kita pakai string)
func getFullNamed() (firstName string, midleName string, lastName string) {
	firstName = "Hariri"
	midleName = "Gifari"
	lastName = "Reptile"

	return
}

func main() {
	// di named return value kita bisa menuliskan seperti ini
	a, b, c := getFullNamed()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

}
