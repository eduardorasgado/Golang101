package main

import "fmt"

func main() {
	var temas [6]string
	setTemas(&temas)
	showMenu(&temas)
}

func setTemas(temas *[6]string) {
	for i := 0; i < 6; i++ {
		output := fmt.Sprintf("tema %d", i)
		temas[i] = output
	}
}

func showMenu(temas *[6]string) {
	fmt.Printf("Seleccione el tema que va a usar: ")
	for i :=0; i < 6;i++ {
		fmt.Printf("%d. %s\n", i, temas[i])
	}
}
