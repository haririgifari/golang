// Struct Method
// 1. struct adalah tipe data seperti tipe data lainnya, dia bisa digunakan sebagai parameter untuk function
// 2. Namun jika kita ingin menambahkan method ke dalam structs, sehingga seakan-akan sebuah struct memiliki function
// 3. Method adalah Function

package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

//kalau mau bikin lagi tinggal
func (a Customer) sayHa() {
	fmt.Println("Hallo kamu sangat tampan", a.Name)
}

func main() {
	var hariri Customer
	hariri.Name = "Hariri Gifari"
	hariri.Address = "Indonesia"
	hariri.Age = 25

	hariri.sayHi("Gifari")
	hariri.sayHa()
}
