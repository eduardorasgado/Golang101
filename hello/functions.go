// Functions in Golang
package main

import "fmt"

const sentence string = "Your name is %s isn't it?\n"

func main() {
	//
	name := getName()
	fmt.Printf(sentence, name)
	
	// call the function and multiple assignments
	a, b, c := getVariables()
	fmt.Print(a, b, c)
	fmt.Print("\n")
	
	var1 := 3
	var2 := 6
	fmt.Print(var1, var2, "\n")
	// easy switching values in variables
	var2, var1 = var1, var2
	fmt.Print(var1, var2, "\n")
	
	// calling a function and passing it a param
	var variable3 int = sum(var1, var2)
	fmt.Print(variable3, "\n")
}

// creating a function
// func name_of_function() type_of_return {}
func getName() string {
	var name string
	fmt.Print("Introduce tu nombre: ")
	// reference to name memory address
	fmt.Scanf("%s", &name)
	
	return name	
}

// returning more than one value
func getVariables() (int, int, int) {
	//
	var number int = 2

	return number, number*2, number*3
}

// function receives values
// type name_variable), (return_type
func sum(a int, b int) (int) {
	//
	return a+b
}