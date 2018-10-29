package main

import "fmt"

// creating a const value
const greetings string = "Hi, how are you?"
// creating a const that can take values as params
const customGreetings string = "Hi %s, what's going on?"

func main() {
	var variable1 string = "Esta es una cadena de texto"
	// podemos hacer que una variable puede ser dinamica
	lastname := "<Dato dinamico, solo puede sufrir modificaciones de tipo configurado al inicio>"
	
	fmt.Printf("%s\n", variable1)

	// cambio del valor de la variable
	variable1 = "cambia la variable1"
	fmt.Printf("%s\n", variable1)
	// usando la variable
	fmt.Printf("%s\n", lastname)

	// way to assign multiple variables
	var (
		a = 1
		b = 2
		c = 3
	)

	fmt.Println(a,b,c)
	
	name := "Eduardo"
	// using a const data
	fmt.Println(greetings)
	// using a const data and introducing a variable into
	fmt.Printf(customGreetings, name)
	fmt.Print("\n")
}