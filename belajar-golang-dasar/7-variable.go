// Variable
// 1. Variable adalah tempat untuk meyimpan data
// 2. Variable digunakan agar kita bisa mengakses data yang sama dimanapun kita mau
// 3. di golang variable hanya bisas menyimpan data yang sama, jika kita ingin menyimpan data yang berrbeda jenis, kita harus membuat beberapa variable
// 4. Untuk membuat variable, kita bisa menggunakan kata kunci var, lalu diikuti dengan nama variable dan tipe datanya
// tidak bisa menggunakan tipe data integer atau boolean (harus string)

package main

import "fmt"

func main() {
	var name string

	name = "Hariri"
	fmt.Println(name)

	name = "Gifari"
	fmt.Println(name)

}

// Tipe data variable
// 1. saat kita membuat variable, maka kita wajib menyebutkan tipe data variable tersebut
// 2. namun jika kita langlung menginisialisasi data pada variable nya, maka kita tidak wajib menyebutkan tipe data variablenya

//contoh tanpa menyebut var name string

package main

import "fmt"

func main() {

	var name = "Hariri" // golang akan langsung tau kalau tipe data ini adalah string
	fmt.Println(name)

	var age = "22" // golang akan langsung tau kalau tipe data ini ada integer
	fmt.Println(age)

}

// kata kunci var
// 1. di golang kata kunci var saat membuat variable tidak lah wajib
// 2. asalkan saat membuat variable kita langsung menginisialisasi datanya
// 3. agar tidak perlu menggunakan kata kunci var, kita perlu menggunakan kata kunci := saat menginisialisasikan data variable tersebut

// contohnya kata kunci var tidak wajib

package main

import "fmt"

func main() {

	country := "Indonesia" // golang akan langsung tau kalau tipe data ini adalah string
	fmt.Println(country)

	country = "Amerika" // untuk deklarasi berikutnya tidak perlu :=, hanya perlu =
	fmt.Println(country)

	age := "22" // golang akan langsung tau kalau tipe data ini ada integer
	fmt.Println(age)

}

// Deklarasi Multiple Variable
// 1. Di golang kita bisa membuat variable secara sekaligus banyak
// 2. code yang dibuat akan lebih bagus dan mudah dibaca

//Contohnya

package main

import "fmt"

func main() {

	var (
		firstName = "Hariri"
		lastName  = "Gifari"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

}
