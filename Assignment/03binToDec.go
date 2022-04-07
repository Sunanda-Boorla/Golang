package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	fmt.Print("Enter Number: ")
	var num string
	fmt.Scanln(&num)
	decimal, err := binaryToDecimal(num)
	if err != nil {
		log.Fatal(err)

	}
	fmt.Println("decimal number is :", decimal)

}
func binaryToDecimal(s string) (int, error) {
	base := 1
	i := len(s) - 1
	decimal := 0
	for i >= 0 {
		if s[i] != '0' && s[i] != '1' {
			return -1, errors.New("not a binary number")
		}
		if s[i] == '1' {
			decimal += base
		}
		base = base * 2
		i--

	}
	return decimal, nil

}
