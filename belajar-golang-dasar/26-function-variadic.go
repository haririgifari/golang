// Variadic Function
// 1. Paarmeter yang bereada diposisi terakhir, memiliki kemampuan dijadikan sebuah varargs
// 2. Varargs artinya datanya bisa menerima lebih dari satu input, atau anggap saja semacam Array.
// 3. Apa bedanya dengan parameter biasa dengan tipe data Array?
//	- Jika parameter tipe Array, kita wajib membuat array terlebih dahulu sebelum mengirimkan ke function
//	- Jika parameter menggunakan varargs, kita bisa langsung mengirim data nya, jika lebih dari satu, cukup gunakan tanda koma

package main

import "fmt"

func sumAll(numbers ...int) int { // cirinya ada ... kalau array []

	total := 0
	for _, number := range numbers {
		total += number // + itu jumlah, kalau - berarti dikurang
	}
	return total
}

func main() {
	total := sumAll(10, 20, 30, 40, 50, 60) // kalau mau kosongin data gausah diisi
	// kalau mau masukin tinggal masukin 10, dst
	fmt.Println(total)

	// Slice parameter
	// 1. Kadang ada kasus dimana kita menggunakan variadic function, namun memiliki variable berupa slice
	// 2. kita bisa menjadikan slice sebagai vararg (variable argument) parameter
	// Contoh :

	slice := []int{10, 20, 30, 40, 50, 60}
	total = sumAll(slice...) //penambahan ... adalah untuk memberitahu kalau ini adalah slice, dan kita ingin menjadikan slice ini menjadi vararg
	fmt.Println(total)

}
