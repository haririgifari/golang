// Function as Parameter
// 1. Function tidak hanya bisa kita simpan sebagai value
// 2. Namun juga bisa kita gunakan sebagai parameter untuk function lain

// Function Type Declarations
// 1. Kadang jika function terlalu panjang, agak ribet untuk menuliskannya didalam parameter
// 2. Type declarations juga bisa digunakan untuk membuat alias function, sehingga akan mempermudah kita menggunakan function sebagai parameter

package main

import "fmt"

//contoh menggunakan type declarations
type Filter func(string) string

// jadi tidak perlu menulis seperti ini
// func sayHelloWithFilter(name string, filter func(string) string)

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello ", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Hariri", spamFilter)
	sayHelloWithFilter("Anjing", spamFilter)

}
