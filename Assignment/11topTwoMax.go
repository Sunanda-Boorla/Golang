package main

import (
	"fmt"
)

func topTwoMax(arr []int) {
	firstMax := arr[0]
	secondMax := -1
	for i := 0; i < len(arr); i++ {
		if arr[i] > firstMax {
			secondMax = firstMax
			firstMax = arr[i]
		} else if arr[i] > secondMax && arr[i] < firstMax {
			secondMax = arr[i]

		}
	}
	if secondMax != -1 {
		fmt.Printf("First max: %d, Second max: %d\n", firstMax, secondMax)

	} else {
		fmt.Println("No two top elements")
	}

}

func main() {
	fmt.Print("Enter no of elements: ")
	var n int
	fmt.Scanln(&n)
	array := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Enter %d element: ", i+1)
		fmt.Scanln(&array[i])
	}
	topTwoMax(array)

}
