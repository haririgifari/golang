// Tipe data map
// 1. Pada Array atau Slice, untuk mengakses data, kita menggunakan index number dimulai dari 0
// 2. Map adalah tip data lain yang berisikan kumpulan data yang sama, namun kita bisa menentukan jenis tipe data index yang akan kita gunakan
// 3. Sederhananya, Map adalah tipe data kumpulan key-value (kata kunci-nilai) dimana kata kuncinya bersifat unik, tidak boleh sama
// 4. Berbeda dengan Array dan Slice, jumlah data yang kita masukkan ke dalam Map boleh sebanyak banyaknya, asalkan kata kunci nya berbeda, jika kita gunakan kata kunci sama, maka secara otomatis data sebelumnya akan diganti dengan data baru

package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Hariri",
		"address": "Jakarta",
	}
	// kalau mau nambah data bisa seperti ini
	person["title"] = "programmer"
	person["age"] = "22"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	// Function Map
	// Operasi						Keterangan
	// len(map)						Untuk mendapatkan jumlah data di map
	// map(key)						Mengambil data di map dengan key
	// map[key] = value				Mengubah data di map dengan key
	// make(map[Typekey]TypeValue)	Membuat map baru
	// delete(map, key)				menghapus data di map dengan key

	// CONTOH Make :

	var book map[string]string = make(map[string]string)
	book["tittle"] = "Belajar Go-Lang"
	book["author"] = "Hariri"
	book["ups"] = "Salah"
	fmt.Println(book) // sebelum di hapus
	fmt.Println(len(book))

	// kalau mau menghapus
	delete(book, "ups")

	fmt.Println(book) // sesudah di hapus
	fmt.Println(len(book))

}
