// Konversi Tipe Data
// 1. Di golang kadang kita butuh melakukan konversi tipe data dari satu tipe ke tipe lain
// 2. Misal kita ingin mengkonversi tipe data int32 ke int63 dan lain lain

package main

import "fmt"

func main() {
	// konversi tipe data 1

	var nilai32 int32 = 100000
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	// konversi tipe data 2 jika ingin mengubah konversi menjadi string

	var name = "Hariri"
	var e byte = name[2] // e itu uint8 alias byte
	var eString = string(e)

	fmt.Println(name)
	fmt.Println(eString)

}
