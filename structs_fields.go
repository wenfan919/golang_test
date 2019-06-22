package main
import "fmt"

type struct1 struct {
	i1  int
	f1  float32
	str string
	stru *struct1
}

func main() {
	ms := new(struct1)
	ms.i1 = 10
	ms.f1 = 15.5
	ms.str= "Chris"

	ms1 := new(struct1)
	ms1.i1 = 100
	ms1.f1 = 150.5
	ms1.str= "Chris 0000"
	ms1.stru  = ms

	ms.stru = ms1



	fmt.Println("stru -----: ", ms.stru)

	fmt.Printf("The int is: %d\n", ms.i1)
	fmt.Printf("The float is: %f\n", ms.f1)
	fmt.Printf("The string is: %s\n", ms.str)
	fmt.Println("stru -----: ",ms)

	fmt.Println("stru i1-----: ",ms.stru.stru.stru.stru.stru.i1)
}





type innerS struct {
	in1 int
	in2 int
}

type outerS struct {
	b    int
	c    float32
	int  // anonymous field
	e innerS //anonymous field
	innerS
}

func main2() {
	outer := new(outerS)
	outer.b = 6
	outer.c = 7.5
	outer.int = 60
	outer.in1 = 5
	outer.in2 = 10

	fmt.Printf("outer.b is: %d\n", outer.b)
	fmt.Printf("outer.c is: %f\n", outer.c)
	fmt.Printf("outer.int is: %d\n", outer.int)
	fmt.Printf("outer.in1 is: %d\n", outer.in1)
	fmt.Printf("outer.in2 is: %d\n", outer.in2)

	// 使用结构体字面量
	outer2 := outerS{6, 7.5, 60, innerS{5, 10},innerS{5, 10}}

	var p * outerS

	p = &outer2


	fmt.Println("outer2 is:", p)
	fmt.Println("outer2 b is:", p.b)
	fmt.Println("outer2 e is:", p.e)
}