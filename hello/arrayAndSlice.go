package main

import "fmt"

func main() {
	/* arrays and slice
	* Arrays: fixed length
	* Slice: dynamic length
	*/
	getArray()
	
	getSlice()
}

func getArray() {
	// we should specify the length for the array
	var array1 [2]string
	// filled values in array
	array1[0] = "value1"
	array1[1] = "value2"
	// printing the array
	fmt.Println(array1)
	
	// other way to create an array
	arr2 := [10]int{1,2,3,4,5,6,7,8,9}
	fmt.Println(arr2)
}

func getSlice() {
	// creating slices: slice is a dynamic container
	var slice1 []string
	// to fill or add values inside an slice we should use append
	slice1 = append(slice1, "Maria", "Pedro", "Mateo", "Juanisimo")
	fmt.Println(slice1)

	// other way to create slices
	var slice2 []int
	slice2 = append(slice2, 1,2,3,4,5,6,7,8,9)
	fmt.Println(slice2)
}