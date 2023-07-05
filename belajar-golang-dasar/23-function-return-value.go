// Function Return Value
// 1. Function bisa mengeembalikan data
// 2. Untuk memberitahu bahwa function mengembalikan data, kiat harus menuliskan tipe data kembalian dari function tersebut
// 3. Jika Function tersebut kita deklarasikan dengan tipe data pengembalian, maka wajib di dalam function nya kita harus mengembalikan data
// 4. Untuk mengembalikan data dari function, kita bisa menggunakan kata kunci return, diikuti dengan datanya

package main

import "fmt"

//							dibawah ini tipe data pengembaliannya
// func getHello(name string) string {
// 	return "Hello " + name
// 	// return itu disimpan dibagian paling bawah, karena setelah return dia akan kembali functionnya, jadi blok kode di bawah return itu tidak akan di eksekusi
// }

// func main() {
// 	result := getHello("Hariri")
// 	fmt.Println(result)
// }

// Contoh menggunakan return yang berbeda beda
func getHello(name string) string {
	if name == "" {
		return "Hello bro"
	} else {
		return "Hello " + name
	}
	//jadi return bisa di pakai berkali kali
	// return itu disimpan dibagian paling bawah, karena setelah return dia akan kembali functionnya, jadi blok kode di bawah return itu tidak akan di eksekusi
}

func main() {
	result := getHello("Hariri")
	fmt.Println(result)

	fmt.Println(getHello(""))
}
