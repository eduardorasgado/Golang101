package main

import "fmt"

func main() {
	// numerical variable types
	var var1 uint;
	var var2 int32;
	var var3 int64;

	var1 = 120
	var2 = 12334
	var3 = 6654464

	value := returningValues(var1, var2, var3)
	fmt.Printf("%d\n", value)
	
	a, b, c := returningValues2(var1, var2, var3)
	fmt.Print(a, b, c, "\n");

	// floating values
	f1, f2, f3 := getFloat()
	fmt.Print(f1, f2, f3, "\n");
}

func getFloat() (float32, float64, float64) {
	// returning floats
	// these conversions always tries to convert the values
	// with the precision it can obtain
	// usually if float32 is converted to float64 it lose precision
	return float32(0.11), float64(0.4534), float64(float32(0.1))
}

func returningValues(var1 uint, var2 int32, var3 int64) int {
	fmt.Print(var1, var2, var3, "\n");
	return 3;
}

func returningValues2(var1 uint, var2 int32, var3 int64) (uint, int32, int64) {
	fmt.Print(var1, var2, var3, "\n");
	return (var1+500),(var2+520),(var3+600);
}