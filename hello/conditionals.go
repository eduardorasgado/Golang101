package main

import "fmt"

func main() {
	//
	ifTest()
	guessNumber()	
}

func ifTest() {
	// if conditional
	// requesting data from user
	var number int;
	fmt.Print("Insert a number from 1 to 10: ")
	fmt.Scanf("%d", &number)

	// checking if number is even or odd
	if number%2 == 0 {
		fmt.Println("An even number.")
	} else {
		fmt.Println("An odd number.")
	}

	// other way to make a conditional
	// here we declare the variable inside the if
	if number2 := 3; number == number2 {
		fmt.Println("The number is equals to 3")
	}
}

func guessNumber() {
	// requesting number from user
	var userNumber int;
	fmt.Print("Insert your number")
	fmt.Scanf("%d", &userNumber)

	// conditional
	// defining a new value; conditional to execute
	if guessNumber :=5; userNumber == guessNumber {
		fmt.Printf("You guessed, it is %d\n", guessNumber)
	} else {
		fmt.Println("Nah! Your wrong")
	}
}