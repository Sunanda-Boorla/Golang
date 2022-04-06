package main

import (
	"fmt"
)

func secondMaximum(arr []int) {
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
		fmt.Printf("second max: %d\n", secondMax)

	} else {
		fmt.Println("No second max")
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
	secondMaximum(array)

}
