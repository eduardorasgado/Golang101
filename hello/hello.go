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
}