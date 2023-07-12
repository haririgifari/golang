// dalam satu package tidak boleh ada function nama yang sama
go env -w GO111MODULE=off
package helper

import "fmt"

func SayHello(name string) {
	fmt.Println("Hello", name)
}
