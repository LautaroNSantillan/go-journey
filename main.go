package main

import "fmt"

func main() {

	//var & consts
	bookingAppName := "The greatest Go booking app ever" //syntactic sugar
	const totalTickets = 50
	var remainingTickets uint = 50

	//array & slices
	bookings := []string{} // no size = slice

	fmt.Printf("bookingAppName is %T, totalTickets is %T, remainingTickets is %T\n", bookingAppName, totalTickets, remainingTickets)

	//println & printf
	fmt.Printf("Welcome to %v, enjoy your stay\n", bookingAppName)
	fmt.Println("You are about to buy tickets")
	fmt.Printf("Total tickets: %v, Available tickets: %v\n", remainingTickets, totalTickets)

	//data types
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	//user input
	fmt.Println("Please enter your first name")
	fmt.Scan(&firstName) //& pointer to where the var is stored in memory

	fmt.Println("Please enter your last name")
	fmt.Scan(&lastName)

	fmt.Println("Please enter your email")
	fmt.Scan(&email)

	fmt.Println("Please enter your amount of desired tickets")
	fmt.Scan(&userTickets)

	remainingTickets -= userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive your order details at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remainig for %v.\n", remainingTickets, bookingAppName)

	fmt.Printf("All bookings: %v\n", bookings)
}
