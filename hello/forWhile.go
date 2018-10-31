package main

import "fmt"

func main() {
	// in go there are not different keywords for while and for loops
	// we have a for for both either
	// normal for
	forTest()
	// for -> while behavior
	forTest2()
	infiniteFor()
}

func forTest() {
	// from 0 to 4
	for i := 0; i < 5; i++ {
		fmt.Println("[",i,"]")
	}
}

func forTest2() {
	// for statement with while behavior
	// no iterator initialized
	// while < is greater than 0
	c := 100
	for c > 0 {
		// each iteration it will decrease
		c -= 10
		fmt.Println("[FOR] while condition behavior, then c is: ", c)
	}
}

func infiniteFor() {
	// infinite way but a bit risky
	// 10,000,000,000 -> for loop in 7 seconds
	// 1,000,000,000 / count up to 20 each one -> 10 seconds(6gb RAM, i7)
	s := 1000000000
	for {
		s -= 1
		for i := 0;i < 20;i++ {
			// Counting up to 20 each iteration over the first loop scope
		}
		//fmt.Printf("%d ", s)
		if s == 0 {
			fmt.Println("\nEnd while loop(infinite for) with a break.")
			// a for loop with while behavior can stop with a break
			break
		}
	}
}