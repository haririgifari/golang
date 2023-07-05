// Anonymous Function
// 1. Sebelumnya setiap membuat function, kita akan selalu memberikan sebuah nama pada function tersebut
// 2. namun kadang ada kalanya lebih mudah membuat function secara langsung di variable atau parameter tanpa harus membuat function terlebih dahulu
// 3. Hal tersebut dinamakan anonymous function, atau function tanpa nama

package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You Are Blooked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "admin"
	}
	registerUser("admin", blacklist)
	registerUser("Hariri", blacklist)

	registerUser("root", func(name string) bool {
		return name == "root"
	})

	registerUser("Hariri", func(name string) bool {
		return name == "root"
	})
}
