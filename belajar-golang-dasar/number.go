// Tipe data Integer 1
// int8 -128 sd 127
// int16 -32768 sd 32767
// int32 -2147483648 sd 2147483648
// int64 -9223372036854775808 sd 9223372036854775807

// Tipe data Integer 2
// uint8 0 sd 255
// uint16 0 sd 65535
// uint32 0 sd 4294967295
// uint64 0 sd 18446744073709551615

// Tipe data Floating Point
// float32 1.18 x 10 (-pangkat 38) sd 3.4 x 10 (pangkat 38)
// float64 2.23 x 10 (-pangkat 308) sd  1.80 x 10 (pangkat 308)
// complex64 (complex numbers with float32 real and imaginary parts) untuk program mtk atau statistik
// complex128 (complex number with float64 real and imaginary parts) untuk program mtk atau statistik

// Alias (nama lain)
// byte  (uint8)
// rune (int32)
// int (minimal int32) tergantung sistem operasinya kalau 32 bit berarti 32 kalau 64 bit berartit 64
// uint (minimal uint32) tergantung sistem operasinya kalau 32 bit berarti 32 kalau 64 bit berartit 64

package main

import "fmt"

func main() {
	fmt.Println("Satu =", 1)
	fmt.Println("Dua =", 2)
	fmt.Println("Tiga Koma Lima =", 3.5)
}
