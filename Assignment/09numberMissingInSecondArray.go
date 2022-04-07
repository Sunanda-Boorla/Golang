package main

import "fmt"

func findMissing(arr1, arr2 []int) {
	len_arr1 := len(arr1)
	len_arr2 := len(arr2)

	myMap := make(map[int]int)
	for i := 0; i < len_arr2; i++ {
		myMap[arr2[i]] = 1

	}
	for i := 0; i < len_arr1; i++ {
		_, ok := myMap[arr1[i]]
		if !ok {
			fmt.Print(arr1[i], " not in second array\n ")

		}

	}
}

func main() {
	fmt.Print("Enter no of elements in array 1: ")
	var n1 int
	fmt.Scanln(&n1)
	array1 := make([]int, n1)
	for i := 0; i < n1; i++ {
		fmt.Printf("Enter %d element: ", i+1)
		fmt.Scanln(&array1[i])
	}
	fmt.Print("Enter no of elements in array 2: ")
	var n2 int
	fmt.Scanln(&n2)
	array2 := make([]int, n2)
	for i := 0; i < n2; i++ {
		fmt.Printf("Enter %d element: ", i+1)
		fmt.Scanln(&array2[i])
	}
	findMissing(array1, array2)
}
