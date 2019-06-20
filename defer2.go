package main

import "fmt"

func main() {
	// fb()

	//var s string

	var mapLit map[string]int
	//var mapCreated map[string]float32
	//var mapAssigned map[string]int

	mapLit = map[string]int{"one": 1, "two": 2}
	//mapCreated := make(map[string]float32)
	//mapAssigned = mapLit

	//mapCreated["key1"] = 4.5
	//mapCreated["key2"] = 3.14159
	//mapAssigned["two"] = 3

	s := last(mapLit)

	fmt.Println("input address4 : " , &s)

	fmt.Println("main :" , s)

}

func last(s map[string]int ) (r map[string]int){

	fmt.Println("input : " , s)

	fmt.Println("input address1 : " , &s)
	//defer change(s)
	defer change2(&s)

	fmt.Println("output : ", s)
	fmt.Println("input address2 : " , &s)
	return  s
}

//func change(s map[string]int){
//
//	s = "last"
//
//	fmt.Println("change : ", s)
//}

func change2(sp *map[string]int){

	//*sp["one"] = 3

	(*sp)["one"] = 3
	fmt.Println("change2 : ", *sp)

	fmt.Println("input address3 : " , sp)
}

