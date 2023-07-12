// If Expression
// 1. if adalah salah satu kata kunci yang digunakan untuk percabangan
// 2. percbangan artinya kita bisa mengeksekusi kode program tertentu ketika suatu kondisi terpenuhi
// 3. hampir di semua bahasa pemrograman mendukung if expression

package main

import "fmt"

func main() {
	var name = "Hariri"

	if name == "Hariri" {
		// jika name hasil boolean di if true maka akan di eksekusi
		// jika name hasil boolean di if false maka tidak akan di eksekusi
		fmt.Println("Hello Hariri")
	} else if name == "Gifari" {
		fmt.Println("Hello Gifari")
	} else if name == "Reptile" {
		fmt.Println("Hello Reptile")
	} else { // menggunakan else jika if false maka yang di print adalah Hi, Boleh kenalan?
		fmt.Println("Hi, Boleh kenalan?")
	} // CONTOH Else Expression
	// Else Expression
	// 1. blok if akan dieksekusi ketika kondisi if bernilai true
	// 2. kadang kita ingin melakukan eksekusi program tertentu jika kondisi if bernilai false
	// 3. hal ini bisa di lakukan menggunakan else expression

	// Else if Expression
	// 1. kadang dalam if, kita butuh membuat beberapa kondisi
	// 2. kasus seperti ini, kita bisa menggunakan Else if Expresssion ( dianta if dan else)

	// IF dengan Short Statement
	// 1. if mendukung short statement sebelum kondisi
	// 2. hal ini sangat cocok untuk membuat statement yang sederhana sebelum melakukan pengecekan terhadap kondisi

	if length := len(name); length > 6 {
		fmt.Println("Nama Terlalu Panjang")
	} else {
		fmt.Println("Nama Sudah Benar")
	}
}
