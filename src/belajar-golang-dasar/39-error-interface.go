// Error Interface
// Go-Lang memiliki interface yang digunakan sebagai kontrak untuk membuat error, nama interface nya adalah error
// contoh nya
// type error interface {
// 	Error() string
// }

// Membuat Error
// 1. Untuk membuat error, kita tidak perlu manual
// 2. Go-Lang sudah menyediakan library untuk membuat helper secara mudah, yang terdapat di package errors ( Package akan kita bahas secara detail di materi tersendiri)

package main

import (
	"errors"
	"fmt"
)

// kode program error interface 1
func Pembagi(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	} else {
		result := nilai / pembagi
		return result, nil
	}
}

func main() {
	var contohError error = errors.New("Ups Error")
	fmt.Println(contohError.Error())

	// contoh menggunakan kode pembagi diatas
	hasil, err := Pembagi(100, 0)
	if err == nil {
		fmt.Println("Hasil", hasil)
	} else {
		fmt.Println("Error", err.Error())
	}
}
