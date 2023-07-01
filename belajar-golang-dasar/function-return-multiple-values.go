// Returning multiple values (bisa lebih dari satu)
// 1. function tidak hanya dapat mengembalikan satu value, tapi juga bisa multiple value
// 2. Untuk memberitahu jika function mengembalikan multiple value, kita harus menulis semua tipe data return value nya di function

// Menghiraukan Return Value
// 1. Multiple return value wajib ditangkap semua value nya
// 2. Jika kita ingin menghiraukan return value tersebut, kita bisa menggunakan tanda _ (garis bawah)
package main

import "fmt"

func getFullName() (string, string, string) { // bisa lebih dari banyak data
	return "Hariri", "Gifari", "Reptile"
}

func main() {
	// contoh menghiraukan return value dengn _
	firstName, _, lastName := getFullName()
	fmt.Println(firstName, lastName)
	fmt.Println(firstName)
	// fmt.Println(midleName)
	fmt.Println(lastName)
}
