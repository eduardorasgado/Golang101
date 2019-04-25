package main

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
