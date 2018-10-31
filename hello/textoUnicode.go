package main

import "fmt"

func main() {
	//
	stringUTF8Japanese := getUnicode();
	fmt.Printf("%s\n", stringUTF8Japanese)

	stringText2 := getTextByIndex(stringUTF8Japanese)
	fmt.Printf("byte type by calling text[i]: %d\n", stringText2)

	// getting a string from user in console
	var yourName string
	fmt.Print("Insert your name: ")
	fmt.Scanf("%s", &yourName)
	
	var t3 string = getTextByIndex2(yourName)
	fmt.Printf("getting string by index and string casting: %s\n", t3)

	// geting length of a string
	tLength := getLengthString(yourName)
	fmt.Printf("%d\n", tLength)
}

func getUnicode() string {
	// handlig text UTF-8
	return "コトリンゴ -「 悲しくてやりきれない 」"
}

// get a character by index will return a byte type value
func getTextByIndex(myText string) byte {
	//
	var textIndex byte = myText[0]
	return textIndex 
}

// we need to make a convertion before returning a value 
// of index string
func getTextByIndex2(myText string) string {
	// making a string casting to a byte type
	var textIndex string = string(myText[0])
	return textIndex
}

func getLengthString(t string) int {
	// returns string length
	return len(t)
}