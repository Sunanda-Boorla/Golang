package main

import (
	"fmt"
)

func areAnagrams(str1, str2 string) bool {
	var count [256]int

	if len(str1) != len(str2) {
		return false

	}
	for i := 0; i < len(str1); i++ {
		count[str1[i]]++
		count[str2[i]]--
	}

	for i := 0; i < 256; i++ {
		if count[i] != 0 {
			return false
		}

	}
	return true

}

func main() {
	var str1 string
	var str2 string
	fmt.Print("Enter String 1: ")
	fmt.Scanln(&str1)
	fmt.Print("Enter String 2: ")
	fmt.Scanln(&str2)
	if areAnagrams(str1, str2) {
		fmt.Printf("%s, %s are Anagrams", str1, str2)
	} else {
		fmt.Printf("%s, %s are not Anagrams", str1, str2)
	}
	fmt.Println()
}
