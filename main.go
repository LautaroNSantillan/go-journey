package main

import "fmt"

func main() {
	fmt.Println("Hello world!")

	//strongly typed - compiled
	var myString string = "This is a string"
	//unused variables wont let the program run, gotta use them
	fmt.Println(myString)

	myString = "This is a new string"

	var declaredString string
	declaredString = "You can declare them then assign them"
	fmt.Println(declaredString)

	var noTypeString = "No type specified"
	fmt.Println(noTypeString)

	var myInt int = 32
	fmt.Println(myInt)
	myInt = +10
	fmt.Println(myInt)

	fmt.Println(myString, myInt)
}
