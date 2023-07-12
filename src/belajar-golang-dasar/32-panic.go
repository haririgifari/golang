// Defer, panic dan Recover

// Panic
// 1. panic function adalah function yang bisa kita gunakan untuk menghentikan program
// 2. panic function biasanya dipanggil ketika terjadi error pada saat program kita berjalan
// 3. saat panic function dipanggil, program akan terhenti, namun defer function tetap akan dieksekusi

package main

import "fmt"

func endApp() {
	fmt.Println("Aplikasi selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("APLIKASI ERROR")
	}
	fmt.Println("Aplikasi Berjalan")
}

func main() {
	// endApp()
	runApp(true) //Aplikasi selesai
	//panic: APLIKASI ERROR
	runApp(false) //Aplikasi selesai
	//Aplikasi berjalan
}
