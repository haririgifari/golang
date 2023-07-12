// Pass by Value
// 1. secara default di Go-Lang semua variable itu di passing by value, bukan by reference
// 2. artinya jika kita mengirim sebuah variable ke dalam function, method atau variable lainn, sebenarnya yang dikirim adalah duplikasi valuenya

// Pointer
// 1. Pointer adalah kemampuan membuat reference ke lokasi data di memory yang sama, tanpa menduplikasi data yang sudah ada
// 2. sederhananya dengan kemampuan pointer, kita bisa membuat pass by reference

// Operator &
// 1. Untuk membuat sebuah variable dengan nilai pointer ke variable yang lain, kita bisa menggunakan operator & diikuti dengan nama variable nya

// Operator *
// 1. Saat kita mengubah variable pointer, maka yang berubah hanya variable tersebut
// 2. Semua variable yang mengacu kedata yang sama tidak akan berubah
// 3. Jika kita ingin mengubah seluruh variable yang mengacu ke data tersebut, kita bisa menggunakan operator *

// Function New
// 1. Sebelumnyaa untuk membuat pointer dengan menggunakan operator &
// 2. Go-lang juga memiliki function new yang bisa digunakan untuk membuat pointer
// 3. Namun function new hanya mengembalikan pointer ke data kosong, artinya tidak ada data awal

// Pointer di Function
// 1. Saat kita membuat parameter di function, secara default adalah pass by value, artinya data akan di copy lalu dikirim ke function tersebut
// 2. Oleh karena itu, jika kita mengubah data di dalam function, data yang asli tidak akan pernah berubah
// 3. Hal ini membuat variable menjaid akan, karena tidak akan bisa diubah
// 4. Namun kadang kita ingin membuat function yang bisa mengubah data asli parameter tersebut
// 5. Untuk melakukan ini, kita juga bisa menggunaka pointer di function
// 6. Untuk menjadikan sebuah parameter sebagai pointer, kita bisa menggunakan operator * di parameternya

package main

import "fmt"

type Address struct {
	City, Porvince, Country string
}

// Contoh Pointer di Function
func ChangeAddressToIndonesia(address *Address) {
	address.Country = "Indonesia"
}
func main() {
	//Pass By Value
	address1 := Address{"Jakarta", "Dki Jakarta", "Indonesia"}
	address2 := address1
	// Operator &
	address3 := &address1
	// Operator *
	address4 := &address1
	*address4 = Address{"Malang", "Jawa Timur", "Indonesia"}

	address2.City = "Bandung"
	*&address2 = Address{"Malang", "Jawa Timur", "Indonesia"}

	//Function NEW
	address5 := new(Address)
	address5.City, address5.Porvince, address5.Country = "Bali", "Denpasar", "Indonesia"

	// Contoh menggunakan Pointer di Function
	alamat := Address{
		City:     "Ajibarang",
		Porvince: "Jawa Tengah",
		Country:  "",
	}
	ChangeAddressToIndonesia(&alamat)
	fmt.Println(alamat)

	fmt.Println(address1) // {Jakarta Dki Jakarta Indonesia}
	fmt.Println(address2) // {Bandung Dki Jakarta Indonesia}
	fmt.Println(address3) // &{Jakarta Dki Jakarta Indonesia}
	fmt.Println(address4) // adanya * semua berubah menjadi {Malang Jawa Timur Indonesia}
	fmt.Println(address5) // {Kosong}
}
