// type declarations
// 1. type declarations adalah keampuan membuat ulang tipe data baru dari tipe data yang sudah ada
// 2. type declarations biasanya digunakan untuk membuat alias terhadap tipe data yang sudah ada dengan tujuan agar lebih mudah dimengerti

// membuat alias untuk tipe data yang sudah ada

package main

import "fmt"

func main() {
	// type no ktp untuk tipe string (no ktp itu alias)
	type NoKTP string
	// contoh pakai boolean
	type Married bool

	var noKtpEko NoKTP = "20202020202020202020"
	var marriedStatus Married = true

	fmt.Println(noKtpEko)
	fmt.Println(marriedStatus)

}
