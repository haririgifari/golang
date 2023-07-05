//Operasi Perbandingan
// 1. Operasi perbandingan adalah operasi untuk membandingkan dua buah data
// 2. Operasi perbandingan adalah operasi yang menghasilkan nilai boolean (benar atau salah)
// 3. Jika hasil operasinya adalah benar, maka nilainya adalah true
// 4. Jika hasil operasinya adalah salah, maka nilainya adalah false

// Operator    Keterangan
//	>			Lebih dari
//	<			Kurang dari
//	>=			lebih dari sama dengan
//	<=			kurang dari sama dengan
//	==			sama dengan
//	!=			tidak sama dengan
// tidak hanya angka yang bisa di bandingkan

package main

import "fmt"

func main() {
	var name1 = "Hariri"
	var name2 = "Gifari"
	var name3 = "HGR"
	var name4 = "HGR"

	var result bool = name1 == name2
	var result1 bool = name3 == name4
	fmt.Println(result)
	fmt.Println(result1)

	// untuk angka / integer

	var value1 = 100
	var value2 = 200

	fmt.Println(value1 > value2)  //false (salah karena lebih kecil v1)
	fmt.Println(value1 < value2)  //true (benar v1 lebih kecil dari v2)
	fmt.Println(value1 == value2) //false (salah karena tidak sama)
	fmt.Println(value1 != value2) //true (karena memang v1 dan v2 berbeda)
}
