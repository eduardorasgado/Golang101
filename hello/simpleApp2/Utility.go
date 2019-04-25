package main

import "fmt"

/**
Devolvemos un conjunto de valores de distintos tipos en una funcion
*/
func returnCouple(value1 int) (float32, int, int) {
	var (
		timesDividedBy2 float32
		times2          int
		times3          int
	)
	timesDividedBy2 = float32(value1) / 2
	times2 = value1 * 2
	times3 = value1 * 3

	return timesDividedBy2, times2, times3
}

/**
Una funcion que toma parametro por referencia, modifica el valor que contiene
esta direccion de memoria
*/
func addDouble(x *int) {
	// el valor dentro de la direccion x o &x es *x y solo modificamos el valor
	// dentro de esta referencia usando
	//doubleX := (*x) * (*x)
	*x += *x
}

/**
Funcion para llamar funciones
*/
func workingWithArrays() {
	// example 1
	exampleArray()
	exampleArray2()
}

/**
Funcion para trabajar con arreglos exclusivamente y con el ciclo for
*/
func exampleArray() {
	var array1 [4]int
	array1[0] = 1
	array1[1] = 2
	array1[2] = 3
	array1[3] = 4

	for i := 0; i < len(array1); i++ {
		fmt.Printf("%d ", array1[i])
	}
	fmt.Println()
}

func exampleArray2() {
	array1 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := 0; i < len(array1); i++ {
		fmt.Printf("%d ", array1[i])
	}
	fmt.Println()
}
