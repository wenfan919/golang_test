package main

import (
	"sync"
	"fmt"
	"time"
)

type info struct {

	mu sync.Mutex

	int
}

func main (){

	t1 := time.Now()
a := new(info)

a.int = 3

update(a, 5)
	update(a, 5)
	update(a, 5)
	update(a, 5)
	update(a, 5)
	update(a, 5)
	update(a, 5)
	update(a, 5)

t := time.Now().Nanosecond() - t1.Nanosecond()

	fmt.Println("sdfdfdf：", a)
	fmt.Println("sdfdfdf：", t)

}

func update(inf *info, s int){

	inf.mu.Lock()

	inf.int = s

	inf.mu.Unlock()

}
