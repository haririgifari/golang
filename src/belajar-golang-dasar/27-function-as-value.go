// Function Value
// 1. Function adalah first class citizen
// 2. Function juga merupakan tipe data, dan bisa disimpan didalam variable

package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func main() {
	goodbye := getGoodBye
	result := goodbye("Hariri")
	fmt.Println(result)
	fmt.Println(getGoodBye("Hariri"))

}
