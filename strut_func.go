package main

import (
	"container/list"
	"fmt"
	"time"
)

type l list.List

func (p l) Iter() {
	// ...
	//for x := range p {
		fmt.Println("x:",1)
	//}
}

func main() {

	ii := IntVector{1, 2, 3}

	fmt.Println(ii.Sum()) // 输出是6

	lst := list.New()
	//lst.pus
	lst.PushBack("fist")
	lst.PushFront(67)

	ll := new(l)
	ll.Iter()
	//for _= range lst.Iter() {
	//}

	//t := new(tt)
	t := tt{time.Now()}
	fmt.Println(t.first3Chars())

}


type IntVector []int

func (v IntVector) Sum() (s int) {
	for _, x := range v {
		s += x
	}
	return
}

type tt struct {
	time.Time
}

func (t tt) first3Chars() string {

	return t.Time.String()
	//return "TEST"
}