// switch expression
// 1. selain if expression, untuk melakukan percabangan kita juga bisa menggunakan switch expression
// 2. switch expression sangat sederhana dibandingkan if
// 3. biasanya switch expression digunakan untuk melakukan pengecekan ke kondisi dalam satu variable

package main

import "fmt"

func main() {
	name := "Hariri"

	switch name {
	case "Hariri":
		fmt.Println("Hello Hariri")
		fmt.Println("Hello Hariri")
	case "Gifari":
		fmt.Println("Hello Gifari")
		fmt.Println("Hello Gifari")
	default:
		fmt.Println("Hi, Boleh Kenalan")
		fmt.Println("Hi, Boleh kenalan")
	}

	// Switch dengan short statement
	// Sama dengan if, switch juga mendukung short statement sebelum variable yang akan di cek kondisinya

	switch length := len(name); length > 6 {
	case true:
		fmt.Println("Nama Terlalu Panjang")
	case false:
		fmt.Println("Nama Sudah Benar")
	}
	// Switch Tanpa Kondisi
	// 1. Kondisi di switch expression tidak wajib
	// 2. Jika kita tidak menggunakan kondisi di switch expression, kita bisa menambahkan kondisi tesebut di setiap casenya

	length := len(name)
	switch {
	case length > 6:
		fmt.Println("Nama Terlalu Panjang")
	case length > 5:
		fmt.Println("Nama Lumayan Panjang")
	default:
		fmt.Println("Nama Sudah benar")
	}
}
