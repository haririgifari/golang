// Closure
// 1. Closure adalah kemampuan sebuah function berinteraksi dengan data disekitarnya dalam scope yang sama
// 2. Harap gunakan fitur closure ini dengan bijak saat kita membuat aplikas

// hati hati menggunakan closure
package main

import "fmt"

func main() {
	name := "Hariri"
	counter := 0

	increment := func() { // variable diatas selalu bisa akses di bawah kurawal
		name := "Gifari" // kalau pakai := tidak akan berubah, kalau pakai = data di atas berubah
		fmt.Println("increment")
		counter++ // jadi counter di atas harusnya 0, jadi ditambah dengan ini
		fmt.Println(name)
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)

}
