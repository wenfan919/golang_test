package main

import (
	"bytes"
	"fmt"
	"sort"
)

func main() {
	//test()

	//test1()

	//test2()

	//test3()

	//test4()

	//test5()

	//test6()
	//test7()

	//test9()

	//fmt.Println(Compare([]byte("232dfdf"), []byte("sfdsf2323232")))

	test10()
}

func test10(){


	var a []int = []int{3,5,6,2,9,6,7,454,576,232,12,6,343,565,7878,565,3434,222222}

	sort.Ints(a)

	fmt.Println(a)

	//sort.

}

func Compare(a, b[]byte) int {
	for i:=0; i < len(a) && i < len(b); i++ {
		switch {
		case a[i] > b[i]:
			return 1
		case a[i] < b[i]:
			return -1
		}
	}
	// 数组的长度可能不同
	switch {
	case len(a) < len(b):
		return -1
	case len(a) > len(b):
		return 1
	}
	return 0 // 数组相等
}


func test9(){
	s := "testhaha"

	fmt.Println(s)


	b := []byte(s)

	fmt.Println(b)
	b[4] = 'w'

	fmt.Println(b)
	//s2 := string(b)

	//fmt.Println(&s2)

	//c :=

}
func test8 (){
	//var c bytes
	//len(c)

	ss := "\u754c\u754c"

	//s := []int32(ss)

	//s := []rune(ss)

	s := []byte(ss)

	for i, c := range s {
		fmt.Printf("%d: %c: %d \n", i, c, len(s))
	}

}
func test7() {

	var b []byte = []byte{'a', 'b', 'c', 'e', 'f'}

	//var d []byte = []byte{'a', 'b','c','e','f'}

	fmt.Println("sl len is :", len(b))
	fmt.Println("sl cap is :", cap(b))
	c := AppendByte(b, b)

	fmt.Println(c)
	fmt.Println(&c)

	fmt.Println("sl len is :", len(c))
	fmt.Println("sl cap is :", cap(c))

}

func AppendByte(slice []byte, data []byte) []byte {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) { // if necessary, reallocate
		// allocate double what's needed, for future growth.
		newSlice := make([]byte, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}

func test6() {

	sl_from := []int{1, 2, 3}
	sl_to := make([]int, 100)

	n := copy(sl_to, sl_from)
	fmt.Println(&sl_to)
	fmt.Printf("Copied %d elements\n", n) // n == 3
	fmt.Println("sl len is :", len(sl_to))
	fmt.Println("sl cap is :", cap(sl_to))
	sl_to = append(sl_to, 4, 5, 6)
	fmt.Println(&sl_to)
	fmt.Println("sl len is :", len(sl_to))
	fmt.Println("sl cap is :", cap(sl_to))

	sl3 := []int{1, 2, 3}
	sl3 = append(sl3, 4, 5, 6)
	fmt.Println(sl3)

	//sl_to = append(sl_to, sl3)

}

func test5() {

	var ss [3][2]int = [3][2]int{{1, 2}, {4, 5}, {8, 9}}

	for row := range ss {
		for column := range ss[row] {
			//ss[row][column] = 1

			fmt.Printf("ss [%d] [%d] = %d \n", row, column, ss[row][column])
		}
	}
}

func test4() {

	var ss []string = []string{"a", "b", "c"}

	for i, s := range ss {

		fmt.Printf("ss %d = %v \n", i, s)
	}

	for _, s := range ss {
		fmt.Printf("ss abc = %v \n", s)
	}

	for i := range ss {
		ss[i] += "123"

		fmt.Printf("ss %d = %v \n", i, ss[i])
	}

}

func test3() {

	i := 0
	var buffer bytes.Buffer
	for {
		if s, ok := getNextString(i); ok { //method getNextString() not shown here
			buffer.WriteString(s)
		} else {
			break
		}
		i++
	}
	fmt.Print(buffer.String(), "\n")
}
func getNextString(i int) (string, bool) {

	if i == 10 {
		return "abc", false
	}
	//return "abc", true
	return "abc", true
}

func test2() {

	//var p *[]int
	//p = new([]int)
	//p := (new([10]int))[0:]

	//p := new([10]int)[0:5]

	var a *[10]int = new([10]int)

	p := a[0:5]

	fmt.Println("  p:", p)
	fmt.Println("len p : ", len(p))
	fmt.Println("cap p : ", cap(p))

	var s []int
	s = make([]int, 1)

	fmt.Println("  s:", s)
	fmt.Println("len s : ", len(s))
	fmt.Println("cap s : ", cap(s))
}

func test1() {

	var x = []int{2, 3, 5, 7, 11}

	y := x[1:3]

	//y := x[1:5]

	fmt.Println(" len x:", len(x))
	fmt.Println(" cap x:", cap(x))

	fmt.Println(" len y:", len(y))
	fmt.Println(" cap y:", cap(y))

	y = y[0:4]

	fmt.Println(" len y:", len(y))
	fmt.Println(" cap y:", cap(y))
	y[2] = 6
	fmt.Println("  y2:", y[2])

	fmt.Println("  x:", x)
}

func test() {

	var arr1 [6]int
	var slice1 []int = arr1[2:5] // item at index 5 not included!

	// load the array with integers: 0,1,2,3,4,5
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}

	// print the slice
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}

	fmt.Printf("The length of arr1 is %d\n", len(arr1))
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	// grow the slice
	slice1 = slice1[0:4]
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	// grow the slice beyond capacity
	//slice1 = slice1[0:7 ] // panic: runtime error: slice bound out of range
}
