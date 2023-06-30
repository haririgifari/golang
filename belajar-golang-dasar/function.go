// Function
// 1. Sebelumnya kita sudah mngenal sebuah function yang wajib dibuat agar program kita bisa berjalan, yaitu function main
// 2. Function adalah sebuah blok kode yang sengaja dibuat dalam program agar bisa digunakan berulang ulang
// 3. Cara membuat function sangat sederhana, hanya dengan menggunakan kata kunci func lalu diikuti dengan nama functionnya dan blok kode isi functionnya
// 4. Setelah membuat function, kita bisa mengeksekusi function tersebut dengan memanggilnya menggunakan kata kunci nama functionnya diikuti tanda kurung buka, kurung tutup

package main

import "fmt"

func sayHello() {
	fmt.Println("Hello World")
}

func main() {
	sayHello()

	// kalau mau eksekusi menggunakan pengulangan(for)
	for i := 0; i < 10; i++ {
		sayHello()
	}
}
