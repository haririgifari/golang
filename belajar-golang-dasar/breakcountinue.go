// Break & Continue
// 1. break & continue adalah kata kunci yang bisa digunakan dalam perulangan
// 2. break digunakan untuk menghentikan seluruh perulangan
// 3. continue adalah digunakan untuk menghentikan perulangan yang berjalan dan langsung melanjutkan ke perulangan selanjutnya

// KODE PROGRAM BREAK
package main

import "fmt"

func main() {
	// for i := 0; i < 10; i++ {
	// 	if i == 5 {
	// 		break
	// 	}
	// 	fmt.Println("perulangan ke", i)
	// }

	// KODE PROGRAM CONTINUE

	// GANJIL
	// for i := 0; i < 10; i++ {
	// 	if i%2 == 0 { // artinya jika menemukan angka genap maka akan di continue
	// 		continue
	// 	}
	// 	fmt.Println("perulangan ke", i) // 1,3,5,7,9
	// }

	// GENAP
	for i := 0; i < 10; i++ {
		if i%2 == 1 { // artinya jika menemukan angka ganjil maka akan di continue
			continue
		}
		fmt.Println("perulangan ke", i) // 0,2,4,6,8
	}

}
