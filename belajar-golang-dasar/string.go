// Tipe data string
// 1. string ada tipe data kumpulan karakter
// 2. jumlah karakter ddi dalam strirng bisas nol sampai tidak terhingga
// 3. Tipe data String di golang direpresantikan dengan kata kunci string
// 4. Nilai data string di golang selalu diawali dengan karekter "petik dua"

package main

import "fmt"

func main() {
	fmt.Println("Hariri")
	fmt.Println("Hariri Gifari")
	fmt.Println("Hariri Gifari Reptile")
}

// Function Untuk String
// len ("string") Menghitung jumlah karakter di String
// "string"[number] mengambil karakter pada posisi yang ditentukan

package main

import "fmt"

func main() {
	fmt.Println(len("Hariri"))
	fmt.Println("Hariri Gifari"[0])
	fmt.Println("Hariri Gifari Reptile"[1])
}
