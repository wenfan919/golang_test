package main

import "fmt"

type Log struct {
	msg string
}

type Custmer struct {

	name string
	log *Log
}

func main(){

	c := &Custmer{"embed test",&Log{"11111"}}

	c.log.add("22222")
	c.l()

}

func (c *Custmer) l(){
	fmt.Println(c.name , c.log.msg)
}

func (lo *Log) add(s string){
	lo.msg += "\n" + s
}
