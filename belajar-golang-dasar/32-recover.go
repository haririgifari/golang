// Defer, panic dan Recover

// Recover
// 1. recover adalah function yang bisa kita gunakan untuk menangkap data panic
// 2. dengan recover proses panic akan terhenti, sehingga program akan tetap berjalan

package main

import "fmt"

func endApp() {
	message := recover()
	if message != nil { // != tidak sama dengan
		fmt.Println("Error dengan message:", message)
	}
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
	// endApp() // aplikasi selesai
	runApp(false) //Error dengan message: APLIKASI ERROR
	//Aplikasi selesai
	//runApp(false) //Aplikasi berjalan
	//Aplikasi selesai
	//fmt.Println("Hariri") kalau pakai recover hariri keluar saat program berjalan
	// tapi kalo tidak pakai recover maka tidak berjalan kebawahnya
}
