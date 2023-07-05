// Defer, panic dan Recover

// Defer
// 1. Defer function adalah function yang bisa kita jadwalkan untuk dieksekusi setelah sebuah function selesai di eksekusi
// 2. Defer function akan selalu dieksekusi walaupun terjadi error di function yang dieksekusi

// Contoh Defer
package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runAppLication(value int) { // value int itu contoh menggunaka deffer
	defer logging() // deffer digunakan selalu diatas
	fmt.Println("Run appLication")
	//menggunaan deffer
	result := 10 / value
	fmt.Println("Result", result)
	logging()
}

func main() {
	runAppLication(2)
}
