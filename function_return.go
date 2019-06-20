package main

import (
	"fmt"
	"strings"
	"log"
	"time"
)

func main1() {
	// make an Add2 function, give it a name p2, and call it:
	p2 := Add2()
	fmt.Printf("Call Add2 for 3 gives: %v\n", p2(3))
	// make a special Adder function, a gets value 2:
	TwoAdder := Adder(2)
	fmt.Printf("The result is: %v\n", TwoAdder(3))
}

func Add2() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

func Adder(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}

func main() {

	//where := func() {
	//	_, file, line, _ := runtime.Caller(1)
	//	log.Printf("%s:%d", file, line)
	//}

	var where = log.Print

	var f = Adder1()

	start := time.Now()
	//longCalculation()



	where()
	fmt.Print(f(1), " - ")
	fmt.Print(f(20), " - ")
	fmt.Print(f(300))

	where()
	addBmp := MakeAddSuffix(".bmp")
	addJpeg := MakeAddSuffix(".jpeg")

	where()
	// returns: file.bmp
	fmt.Println("bmp:", addBmp("file") )
	//addJpeg("file") // returns: file.jpeg
	fmt.Println("jpeg:", addJpeg("file") )
	where()

	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)

}

func Adder1() func(int) int {
	var x int
	return func(delta int) int {
		fmt.Println("delta:" , delta)
		x += delta
		fmt.Println("x:" , x)
		return x
	}
}


func MakeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}