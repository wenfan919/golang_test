package main

import "fmt"

func main() {
	// fb()

	//var s string

	s := last("test")

	fmt.Println("input address4 : " , &s)

	fmt.Println("main :" , s)

}

func last(s string ) (r string){
	fmt.Println("input : " , s)

	fmt.Println("input address1 : " , &s)
	//defer change(s)
	defer change2(&s)

	fmt.Println("output : ", s)
	fmt.Println("input address2 : " , &s)
	return  s
}

func change(s string){

	s = "last"

	fmt.Println("change : ", s)
}

func change2(sp *string){

	*sp = "last2"

	fmt.Println("change2 : ", *sp)

	fmt.Println("input address3 : " , sp)
}


func fb() {

	defer e(m(s("b")))
	fmt.Println("in b")
	fa()
}

func fa() {
	defer e(m(s("a")))
	fmt.Println("in a")
}

func s(s string) (r string) {
	fmt.Printf("function start with %q", s)
	fmt.Println()
	return s
}

func e(e string) {

	fmt.Printf("function end with %q", e)
	fmt.Println()
}

func m(m string) (r string) {
	fmt.Printf("function middle with %q", m)
	fmt.Println()
	return m
}
