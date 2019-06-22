package main

import ("fmt"
//. "./factory"
. "./factory"
	"unsafe"
)
//import "fac"


func main(){

	f := NewFile(10, "./test.txt")

	fmt.Println("f: ", f)

	s := unsafe.Sizeof(f)

	fmt.Println("s: ", s)

}