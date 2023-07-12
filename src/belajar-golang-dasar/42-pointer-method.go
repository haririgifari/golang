// Pointer di Method
// 1. Walaupun Method akan menempel di struct, tapi sebenarnya data struct yang di akses di dalam method adalah pass by value
// 2. Sangat direkomendasikan menggunakan Pointer di method, sehingga tidak boros memory karena harus selalu di duplikasi ketika memanggil method

package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() { //tambahkan *diparameter Man
	man.Name = "Mr. " + man.Name
	fmt.Println("Di Method", man.Name) // Di Method Mr. hariri
}

func main() {
	hariri := Man{"hariri"}
	hariri.Married()

	fmt.Println(hariri.Name) // hariri
}
