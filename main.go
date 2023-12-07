package main

import (
	"fmt"
	"reflect"
)

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

	fmt.Println(reflect.TypeOf(myInt))

	var myFloat float64 = 6.9
	fmt.Println(myFloat + float64(myInt)) //casting

	var myBool bool = true
	fmt.Println(myBool)

	//to declare and initialize the variable in a single line use :=
	initializedString := ":="
	fmt.Println(initializedString)

	//cons
	const myConst = "This is a const" //no need to use it right away

	//control-flow
	if myInt == 42 {
		fmt.Println("int = 42")
		//else if also available
	} else {
		fmt.Println("something else")
	}

	//data structures

}
