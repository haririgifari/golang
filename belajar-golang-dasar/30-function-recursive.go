// Recursive Function
// 1. Recursive Function adalah function yang memanggil function dirinya sendiri
// 2. Kadang dalam pekerjaan, kita sering menemui kasus dimana menggunakan recursive function lebih mudah dibandingkan tidak menggunakan recursive function
// 3. Contoh kasus yang lebih mudah diselesaikan menggunakan recursive adalah factorial

// Kode program Factorial For Loop
package main

import "fmt"

func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i // result x i
	}
	return result
}

// Menggunakan Kode program factorial recursive
func factorRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorRecursive(value-1)
	}
}

func main() {
	loop := factorialLoop(10)
	fmt.Println(loop)
	// fmt.Println(5 * 4 * 3 * 2 * 1) // ini cara manual

	// kode yang menggunakan factorial recrusive
	recursive := factorRecursive(10)
	fmt.Println(recursive)
}
