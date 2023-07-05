// For Loops(Perulangan)
// 1. Dalam bahasa pemrograman, biasanya ada fitur yang bernama perulangan
// 2. Salah satu fitur perulangan adalah for loops

package main

import (
	"fmt"
)

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke", counter)
		counter++ // kalau tidak ada counter++ maka dia akan perulangan terus sampai ketemu false
	}

	// for dengan statement
	// 1. dalam for, kita bisa menambahkan statement, dimana terdapat 2 statement yang bisa tambahkan di for
	// 2. init statement yaitu statement sebelum for dieksekusi
	// 3. post statement yaitu statement yang akan selalu dieksekusi di akhir tiap perulangan

	//	init statement		kondisi		 post statement
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke", counter)
	}

	// kalau mau akses data array atau slice secara manual

	slice := []string{"Hariri", "Gifari", "Reptile"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	//for range (cara cepat mengambil sebuah data didalam array, slice dan map)
	// 1. for bisa digunakan untuk melakukan iterasi terhadap semua data collection
	// 2. data collection contohnya Array, Slice dan Map

	for i, name := range slice {
		fmt.Println("Index", i, "=", name)
	}
	// ingin println data aja (name) tanpa index
	for _, name := range slice {
		fmt.Println(name)
	}
	// untuk MAP
	person := make(map[string]string)
	person["name"] = "Hariri"
	person["title"] = "Gifari"

	// kalau MAP berupa key, kalau array atau slice berupa index (i)
	for key, name := range person {
		fmt.Println(key, "=", name)
	}

}
