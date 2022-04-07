package main

import (
	"fmt"
	"sort"
)

func main() {
	var employee1 = map[string]int{
		"Mark":  40,
		"Sandy": 50,
		"Rocky": 60,
		"Rajiv": 70,
		"Kate":  80}
	fmt.Println(employee1)

	var employee = make(map[string]int) //Empty map
	fmt.Printf("Length of employee: %d\n", len(employee))
	employee["Vicky"] = 10
	employee["Nick"] = 20
	employee["John"] = 30
	fmt.Println(employee)
	fmt.Printf("Length of employee after adding : %d\n ", len(employee))

	//ACCESSING ITEMS
	fmt.Println(employee["Vicky"])

	//Updating value
	employee["Nick"] = 50
	fmt.Println(employee)

	//Delete items
	delete(employee, "Nick")
	fmt.Println(employee["Nick"]) //accessing deleted item gives 0
	fmt.Println(employee)

	//Iterating over map employee1

	for key, value := range employee1 {
		fmt.Println("Key: ", key, " => ", "Value: ", value)
	}

	//Clear all items
	/*for k := range employee {
		delete(employee, k)

	}
	fmt.Println(employee)
	employee=make(map[string]int)*/

	//sort keys
	fmt.Println("Sorting Map Keys")
	keys := make([]string, 0, len(employee1))

	for k := range employee1 {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k, employee1[k])
	}
	//sort values
	fmt.Println("Sorting Map Values")
	values := make([]int, 0, len(employee1))

	for _, v := range employee1 {
		values = append(values, v)
	}
	sort.Ints(values)
	for _, v := range values {
		fmt.Println(v)
	}

	//Merge 2 maps
	for k, v := range employee {
		employee1[k] = v
	}
	fmt.Println("Map after merging: ", employee1)

}
