package main

import "fmt"

type first struct {
	base1 string
}

type second struct {
	base2 string
}

func (f first) printBase1() string {
	return f.base1
}

func (s second) printBase2() string {
	return s.base2

}

type child struct {
	first
	second
}

func main() {
	c1 := child{
		first{
			base1: "I am in base 1",
		},
		second{
			base2: "I am in base 2",
		},
	}
	fmt.Println(c1.printBase1())
	fmt.Println(c1.printBase2())
}
