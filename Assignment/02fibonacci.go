package main

import "fmt"

func main() {
	fmt.Print("Enter Number: ")
	var num int
	fmt.Scanln(&num)
	printFibonacciSeries(num)
}
func printFibonacciSeries(num int) {
	firstnum := 0
	secondnum := 1
	fmt.Printf("%v %v ", firstnum, secondnum)
	for i := 2; i < num; i++ {

		thirdnum := firstnum + secondnum
		fmt.Printf("%v ", thirdnum)
		firstnum = secondnum
		secondnum = thirdnum
	}
	fmt.Println()
}
