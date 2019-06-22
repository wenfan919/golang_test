package main

import (
	"fmt"
	"unsafe"
)

type Foo map[string]string
type Bar struct {
	thingOne string
	thingTwo int
}

func main() {
	// OK
	y := new(Bar)
	(*y).thingOne = "hello"
	(*y).thingTwo = 1

	fmt.Println("y:", unsafe.Sizeof(*y))
	// NOT OK
	//z := make(Bar) // 编译错误：cannot make type Bar
	//(*z).thingOne = "hello"
	//(*z).thingTwo = 1

	// OK
	x := make(Foo)
	x["x"] = "goodbye"
	x["y"] = "world"

	fmt.Println("y:", x)
	// NOT OK
	u := new(Foo)

	//u = make(u)


	fmt.Println("y:", u)

	//(*u)["x"] = "goodbye" // 运行时错误!! panic: assignment to entry in nil map
	//(*u)["y"] = "world"
}
