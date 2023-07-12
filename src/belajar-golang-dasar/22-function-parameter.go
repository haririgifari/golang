// Function Parameter
// 1. Saat membuat function, kadang-kadang kita membutuhkan data dari luar atau kita sebut parameter
// 2. kita bisa menambahkan parameter di function, bisa lebih dari satu
// 3. parameter tidaklah wajib, jika kita bisa membuat function tanpa parameter seperti sebelumnya yang sudah kita buat
// 4. Namun jika kita menambahkan parameter di function, maka ketika memanggil function tersebut, kita wajib memasukkan data ke parameternya

package main

import "fmt"

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func main() {
	sayHelloTo("Hariri", "Gifari")
	sayHelloTo("Ria", "Oktarina")
}
