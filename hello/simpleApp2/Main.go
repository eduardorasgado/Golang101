package main

/**
Importando librerias
*/
import (
	"fmt"
	"math"
)

/**
Declaracion de constantes en un solo bloque
*/
const (
	ALPHA float32 = 32.122
	GAMMA float32 = 45.789
	BETA  float32 = 122.0012
)

func main() {
	// coordenadas en el espacio
	x := 45
	y := 55
	z := 75
	// otras variables
	var (
		abgSolution  float64
		dividedValue float32
		valueby2     int
		valueBy3     int
	)

	/**
	Comentario de bloque o de varias lineas:
		- linea 1
		- linea 2
		- linea 3
	*/
	showSomething()
	showVectorLength(x, y, z)
	abgSolution = getSimpleSolution(ALPHA, GAMMA, BETA)
	fmt.Printf("Solution is: %f\n", abgSolution)

	dividedValue, valueby2, valueBy3 = returnCouple(x)

	fmt.Printf("values, divided: %f, double: %d, triple: %d\n",
		dividedValue, valueby2, valueBy3)

	fmt.Printf("value in main: %d\n", valueBy3)
	addDouble(&valueBy3)
	fmt.Printf("value in main: %d\n", valueBy3)
}

/**
Funcion para imprimir algo en pantalla
*/
func showSomething() {
	fmt.Printf("This is a formatter\n")
}

/**
Funcion que toma parametros para resolver una norma
*/
func showVectorLength(x int, y int, z int) {
	var length float64
	var floatAddition float64

	// casting
	floatAddition = float64((x * x) + (y * y) + (z * z))
	length = math.Sqrt(floatAddition)

	fmt.Printf("longitud del vector conformador por los puntos"+
		"x: %d, y: %d, z: %d, es: %f\n", x, y, z, length)
}

/**
Funcion que calcula una salida y esta a su vez es devuelta
*/
func getSimpleSolution(x float32, y float32, z float32) float64 {
	var solution float64
	solution = float64(1/(x*y*z)) * math.Exp(float64(x/10))

	return solution
}
