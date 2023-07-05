// Tipe data slice ()
// 1. Tipe data slice adalah potongan dari data Array
// 2. Slice mirip dengan Array, yng membedakan adalah ukuran Slice bisas berubah
// 3. Slice dan Array selalu terkoneksi, dimana slice adalah data yang mengakses sebagian atau seluruh data di Array

// Detail Tipe Data Slice
// 1. Tipe data Slice memiliki 3 data, yaitu pointer, length dan capacity
// 2. Pointer adalah penunjuk data pertama di array pada slice (cara membuat slice itu array dulu baru slice)
// 3. Length adalah panjang dari slice
// 4. Capacity adalah kapasitas dari slice, dimana length tidak boleh lebih dari capacity

// Membuat Slice dari Array
// Membuat Slice			Keterangan
// array[low:high]			Membuat slice dari array dimulai index low sampai index sebelum high
// array[low:]				Membuat slice dari array dimulai index low sampai index akhir di array
// array[:high]				Membuat slice dari array dimulai index 0 sampai index sebelum high
// array[:]					Membuat slice dari array dimulai index 0 sampai index akhir di array

// slice dan array menggunakan bulan 0 = januari - 11 = Desember Maka Slice dari :
// array[4:7]	Slice 1 : pointer = 4 (berarti bulan mei), length = 3 ( dari 4,5 dan 6 (7 tidak termasuk)), capacity = 8 (data dari mei-desember)
// array[6:9]	Slice 2 : pointer = 6 (berarti bulan juli), length = 3 (dari 6,7 dan 8 (9 tidak termasuk)), capacity = 6 (data dari juli-desember)

// Function Slice
// Operasi								Keterangan
// len(slice)							Untuk mendapatkan panjang
// cap(slice)							Untuk mendapatkan Kapasitas
// append(slice,data)					Membuat slice baru dengan menambahkan data ke posisi terakhir slice, jika kapasitas sudah penuh, maka akan membuat array baru
// make([]TypeData, length, Capacity)	Membuat slice baru
// copy(destination, source)			Menyalin slice dari source ke destination

// contoh
package main

import "fmt"

func main() {
	var months = [...]string{ //jika tau capasitasnya berapa bisa langsung masukan kalau tidak bisa isi ...
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}
	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	//hati hati saat mengoprasikan slice dan array
	// months[5] = "Diubah"
	// fmt.Println(slice1)

	// slice1[0] = "Mei lagi"
	// fmt.Println(months)

	var slice2 = months[2:4]
	fmt.Println(slice2) // [maret - mei]

	//append
	var slice3 = append(slice2, "hariri")
	fmt.Println(slice3) // [hasilnya november desember hariri] karena slice dua sudah full sampai desember jadi membuat array baru

	slice3[1] = "Bukan Desember"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(months) //karena tidak menggunakan array baru maka jawabannya [Januari Februari Maret Bukan Desember hariri Juni Juli Agustus September Oktober November Desember]

	//make
	newSlice := make([]string, 2, 5)

	newSlice[0] = "Hariri"
	newSlice[1] = "Gifari"

	fmt.Println(newSlice)      // [Hariri Gifari]
	fmt.Println(len(newSlice)) // 2
	fmt.Println(cap(newSlice)) // 5

	//copy mindahin data
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice) // [Hariri Gifari]

	// kalau lenght cuma 1 seperti di bawah, maka hasilnya cuma 1 / datanya terpotong
	copySlice1 := make([]string, 1, cap(newSlice))
	copy(copySlice1, newSlice)
	fmt.Println(copySlice1) // [Hariri]

	//Hati Hati saat membuat Array
	//Saat membuat array, kita harus berhati-hati jika salah, maka yang kita buat bukanlah array melainkan slice
	// Contoh :

	iniArray := [5]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(iniArray) //[1 2 3 4 5]
	fmt.Println(iniSlice) //[1 2 3 4 5]
	//hasilnya sama tapi array [masukan angka capacity] dan slice [kosong]
}
