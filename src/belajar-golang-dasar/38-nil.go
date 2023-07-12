// Nil
// 1. Biasanya didalam bahasa pemrograman lain, object yang belum diinisialisasi maka secara otomatis nilainya adalah null atau nil
// 2. berbeda dengan Go-Lang, di Go-Lang saat kita buat variable dengan tipe data tertentu, maka secara otomatis akan dibuatkan default value nya
// 3. Namun di Go-Lang ada data nil, yaitu data kosong
// 4. Nil sendiri hanya bisa digunakan di beberapa tipe data, seperti interface, function, map, slice, pointer dan channel

package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	person := NewMap("Hariri")
	fmt.Println(person)

	// atau
	var value map[string]string = NewMap("Ria")

	if value == nil {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(value)
	}
}
