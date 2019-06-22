package factory

import "fmt"

type File struct {
	fd      int     // 文件描述符
	name    string  // 文件名
}



func NewFile(fd int, name string) *File {
	if fd < 0 {
		return nil
	}

	return &File{fd-1, name}
}

func main(){

	f := NewFile(10, "./test.txt")

	fmt.Println("f: ", f)
}