package main

import "fmt"

func main() {

	//var & consts
	var bookingAppName = "The greatest Go booking app ever"
	const totalTickets = 50
	var remainingTickets = 50

	fmt.Printf("bookingAppName is %T, totalTickets is %T, remainingTickets is %T\n", bookingAppName, totalTickets, remainingTickets)

	//println & printf
	fmt.Printf("Welcome to %v, enjoy your stay\n", bookingAppName)
	fmt.Println("You are about to buy tickets")
	fmt.Printf("Total tickets: %v, Available tickets: %v\n", remainingTickets, totalTickets)

	//data types
	var userName string
	var userTickets int

	fmt.Printf("User %v has booked %v tickets", userName, userTickets)
}
