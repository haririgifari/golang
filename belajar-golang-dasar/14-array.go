// Tipe data Array
// 1. Array adalah tipe data yyang berisikan kumpulan data dengan tipe yang ssama
// 2. Saat membuat array, kita perlu menentukan jumlah data yang bisa ditampung oleh Array tersebut
// 3. Daya tampung Array tidak bisa bertambah setelah Array dibuat

// Index di Array
// Index		Data
//	0			Hariri
//	1			Gifari
//	2			Reptile
// Index dimulai dari 0, contoh ada dat hariri, gifari, reptile, jika mau ambil data index hariri maka indexnya adalah 0 dst

package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Hariri"
	names[1] = "Gifari"
	names[2] = "Reptile"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// Membuat Array Secara langsung
	// Di Go-lang kita bisa membuat array secar langsung saat deklarasi variable

	var values = [3]int{ // langsung
		90,
		95,
		80,
	}
	fmt.Println(values)

	var name = [3]string{ // langsung
		"Hariri",
		"Gifari",
		"Reptile",
	}
	fmt.Println(name)

	// kalau mau satu satu
	// fmt.Println(values[0])
	// fmt.Println(values[1])
	// fmt.Println(values[2])

	// fmt.Println(name[0])
	// fmt.Println(name[1])
	// fmt.Println(name[2])

	// Function Array
	// Operasi					Keterangan
	// len(array)				Untuk mendapatkan panjang Array
	// array[index]				Mendapatkan data di posisi index
	// array[index] = value		Mengubah data di posisi index

	// contoh
	var lagi [10]string

	fmt.Println(len(lagi)) // hasailnya 10 karena arraynya 10 walaupun belum ada datanya (stringnya)
}
