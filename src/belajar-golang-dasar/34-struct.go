// Struct
// 1. Struct adalah sebuah template data yang digunakan untuk menggabungkan nol atau lebih tipe data lainnya dalam satu kesatuan
// 2. struct biasanya representasi data dalam program aplikasi yang kita buat
// 3. data di struct disimpan dalam field
// 4. sederhananya struct adalah kumpulan dari field

package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

// Membuat data Struct
// 1. struct adalah template data atau prototype data
// 2. struct tidak bisa langsung digunakan
// 3. namun kita bisa membuat data/object dari struct yang telah kita buat

func main() {
	var hariri Customer
	hariri.Name = "Hariri Gifari"
	hariri.Address = "Indonesia"
	hariri.Age = 25

	fmt.Println(hariri.Name)
	fmt.Println(hariri.Address)
	fmt.Println(hariri.Age)

	// Struct literals
	// Sebelumnya kita telah membuat data dari struct, namun sebenarnya ada banyak cara yang bisa kita gunakan untuuk membuat data dari struct

	Gifari := Customer{
		Name:    "Gifari",
		Address: "Jakarta",
		Age:     25,
	}
	fmt.Println(Gifari)
	// Di saran kan menggunakan deklarasi antara dua diatas
	// atau bisa juga
	HGR := Customer{"HGR", "Jakarta", 25} // ini tetap mengacu kepada type diatas, jadi jika kelebihan maka dia akan error
	fmt.Println(HGR)
}
