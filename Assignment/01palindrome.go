package main

import "fmt"

func main() {
	fmt.Print("Enter String: ")
	var str string
	fmt.Scanln(&str)
	if isPalin(str) {
		fmt.Println(str, "is Palindrome")
	} else {
		fmt.Println(str, "is not a Palindrome")
	}
}

func isPalin(str string) bool {
	i := 0
	for ; i < len(str); i++ {
		j := len(str) - 1 - i
		if str[i] != str[j] {
			return false
		}

	}
	return true
}
