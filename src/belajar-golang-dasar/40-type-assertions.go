// Type Assertions
// 1. Type Assertions merupakan kemampuan merubah tipe data menjadi tipe data yang diinginkan
// 2. Fitur ini sering sekali digunakan ketika kita bertemu dengan data interface kosong

package main

import "fmt"

func random() interface{} {
	return true // jika salah bisa terjadi panic
}

func main() {
	// result := random()
	// // resultString := result.(string)
	// // fmt.Println(resultString)

	// Type Assertions Menggunakan Switch
	// 1. Saat salah menggunakan type assertions, maka bisa berakibat terjadi panic di aplikasi kita
	// 2. Jika panic dan tidak ter-recover, maka otomaatis program kita akan mati
	// 3. Agar lebih aman, sebaiknya kita menggunakan switch expression untuk melakukan type assertions
	// contoh

	result := random()
	switch value := result.(type) {
	case string:
		fmt.Println("value", value, "is string")
	case int:
		fmt.Println("value", value, "is int")
	case bool:
		fmt.Println("value", value, "is boolean")
	default:
		fmt.Println("Unknown")
	}

}
