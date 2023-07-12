// Operator		Keterangan
// 	+			Penjumlahan
//  -			Pengurangan
//  *			Perkalian
//  /			Pembagian
//  %			Sisa Pembagian

package main

import "fmt"

func main() {
	// atau bisa langsung
	var result = 10 + 10
	fmt.Println(result)

	//atau bikin dulu variable untuk menampung datanya
	var a = 10
	var b = 10
	var c = a + b
	var d = a - b
	var e = a * b
	var f = a / b
	var g = a % b

	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

	// Augmented Assignments (untuk mempersingkat kalau mau melakukan penambahan terhadap variable dia sendiri)
	// Operasi Matematika	(singkat >)	Augmented Assignments
	//		a = a + 10						a += 10
	// 		a = a - 10						a -= 10
	//		a = a * 10						a *= 10
	//		a = a / 10						a /= 10
	//		a = a % 10						a %= 10

	// contoh
	var i = 10
	i += 10 // artinya i = i + 10

	fmt.Println(i)

	// Unary Operator (digunakan untuk mempersingkat)
	// Operator		(Singkatan >)	Keterangan
	//	++							a = a + 1
	//	--							a = a - 1
	//	-							negativ (untuk menandakan variablenya negatif tambahkan - di depannya)
	//	+							positive
	// 	!							Boolean kebalikan

	// contoh

	var h = 22
	h++ // h = h + 1

	fmt.Println(h)

	var negative = -100
	var positive = +100 // tidak perlu tambah + untuk positif
	fmt.Println(negative)
	fmt.Println(positive)
}
