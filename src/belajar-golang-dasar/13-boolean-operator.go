// Operator Boolean
// Operator			Keterangan
// &&				Dan			(operasi untuk dua boolean)
// ||				atau		(operasi untuk dua boolean)
// !				kebalikan	(operasi untuk satu boolean)

// Operasi &&
//	Nila1	Operator	Nilai2	Hasil
//	true	&&			true	true
//	true	&&			false	false
//	false	&&			true	false
//	false	&&			false	false

// Operasi || (atau)
//	Nila1	Operator	Nilai2	Hasil
//	true	||			true	true
//	true	||			false	true
//	false	||			true	true
//	false	||			false	false

// Operasi !
// Operator		Nilai2		Hasil
// !			true		false
// !			false		true

package main

import "fmt"

func main() {
	var ujian = 90
	var absensi = 80

	var lulusUjian = ujian > 80
	var lulusAbsensi = absensi > 80
	fmt.Println(lulusUjian)
	fmt.Println(lulusAbsensi)

	var lulus = lulusUjian && lulusAbsensi
	fmt.Println(lulus)

	//kenyataannya penulisannya langsung contoh :
	fmt.Println(ujian >= 80 && absensi >= 80)
	fmt.Println(ujian > 80 || absensi > 80) // kenapa dia true karena seperti rumus di atas true-false = true

}
