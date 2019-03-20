package main

import (
	"fmt"
)

func main() {
	fmt.Println(AllocateVar(5))

	fmt.Println(AllocateVar(7))

}
func AllocateMake(b int) []byte {

	arr := make([]byte, b)
	return arr
}
func AllocateVar(b int) []byte {

	var arr []byte = make([]byte, b)
	return arr
}
