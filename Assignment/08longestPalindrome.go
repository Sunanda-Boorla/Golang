package main

import "fmt"

func longestPalindrome(s string) {
	n := len(s)
	if n <= 1 {
		fmt.Println(s)
	}
	maxLen := 1
	start := 0
	for i := 0; i < n; i++ {
		low := i - 1
		high := i + 1
		for high < n && s[high] == s[i] {
			high++
		}
		for low >= 0 && s[low] == s[i] {
			low--
		}
		for low >= 0 && high < n && s[low] == s[high] {
			low--
			high++
		}
		length := high - low - 1
		if maxLen < length {
			maxLen = length
			start = low + 1
		}

	}
	fmt.Println("Longest Palindrome Substring is: ", s[start:start+maxLen])

}

func main() {
	fmt.Print("Enter String: ")
	var str string
	fmt.Scanln(&str)
	longestPalindrome(str)
}
