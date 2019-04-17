package main

import "fmt"

func main() {
	// creamos una variable de tipo string
	var name string
	// Println te permite imprimir y dar un salto de linea al final
	fmt.Println("Ingresa tu nombre")
	// tomamos los datos del user
	fmt.Scanf("%s", &name)
	// concatenamos con los datos del user
	fmt.Printf("Hola Go! :B, te saluda %s\n", name)
	// para poder correr el codigo
	// go run hello.go
	var x int = giveMeSomething(23)
	x = x + 10
	fmt.Printf("Hello, yo tengo el numero: %d", x)
}

func giveMeSomething(x int) int {
	fmt.Printf("This is something: %d\n", x)
	var y int = x + 10
	return y
}
