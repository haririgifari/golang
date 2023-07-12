// Interface
// 1. Interface adalah tipe data abstract, dia tidak memiliki implementasi langsung
// 2. seebuah interface berisikan definisi-definisi method
// 3. Biasanya interface digunakan sebagai kontrak

package main

import "fmt"

type HasName interface {
	GetName() string //interface harus ada GetName dan dibungkus dengan string
}

func sayHay(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

// Implementasi interface
// 1. Setiap tipe data yang sesuuai dengan kontrak interface, secaraotomatis dianggap sebagai interface tersebut
// 2. sehingga kita tidak perlu mengimplementasikan interface secara manual
// 3. Hal ini agak berbeda dengan bahasa pemrograman lain yang ketika membuat interface, kita harus menyebutkan secara eksplisit akan menggunakan interface mana

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

// Kode program implementasi interface ke 2
// contoh
type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	var Gifari Person
	Gifari.Name = "Gifari"
	sayHay(Gifari)
	// contoh penulisan ke 2
	cat := Animal{
		Name: "push",
	}
	sayHay(cat)
}
