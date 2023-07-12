// Import
// 1. secara standar, file Go-Lang hanya bisa mengakses file Go-Lang lainnta yang berada dalam package yang sama
// 2. Jika kita ingin mengakses file Go-Lang yang berada diluar package, maka kita bisa menggunakan import
go env -w GO111MODULE=off
package main

import (
	"belajar-golang-dasar/helper"
	"fmt"
)

func main() {
	helper.SayHello("Hariri")
	fmt.Println(result)
}
