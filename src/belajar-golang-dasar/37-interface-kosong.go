// Interface Kosong
// 1. Go-lang bukanlah bahasa pemrograman yang berorientasi objek
// 2. Biasanya dalam pemrograman berorientasi objek, ada satu data parent dii puncak yang bisa dianggap sebagai semua implementasi data yang ada di bahasa permrograman tersebut
// 3. Contohh di java ada java.lang.Object
// 4. Untuk menangani kasus seperti ini, di Go-lang kita bisa menggunakan interface kosong
// 5. Interface kosong adalah interface yang tidak memiiliki deklarasi method satupun, hal ini membuat secara otoomatis semua tipe data akan menjadi implementasi nya

// Penggunaan interface kosong di Go-Lang
// 1. Ada banyak contoh penggunaan interface kosong di Go-lang, Seperti :
//   - fmt.Println(a ...interface{})
//	 - panic(v interface{})
//	 - recovery() interface{}
//	 - dan lain lain

package main

import "fmt"

func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "ups"
	}
}

func main() {
	var data interface{} = Ups(3)
	fmt.Println(data)
}
