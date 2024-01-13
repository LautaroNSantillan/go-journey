package main

import (
	"fmt"
	"strings"
)

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

	//for: only loop
	for {
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

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidAmount := userTickets > 0 && userTickets <= remainingTickets

		if isValidAmount && isValidName && isValidEmail {

			remainingTickets -= userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive your order details at %v.\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remainig for %v.\n", remainingTickets, bookingAppName)

			firstNamesSlice := []string{}
			for _, element := range bookings { //range iterates // _ : blank identifier
				var names = strings.Fields(element) //split by space
				firstNamesSlice = append(firstNamesSlice, names[0])
			}
			fmt.Printf("All bookings: %v\n", firstNamesSlice)

			if remainingTickets == 0 {
				fmt.Printf("No more tickets available.\n")
				break
			}
		} else {
			if !isValidName {
				fmt.Printf("Your name is too short")
			}
			if !isValidEmail {
				fmt.Printf("Your email is invalid")

			}
			if !isValidAmount {
				fmt.Printf("Invalid ticket number")

			}
			//continue //next iteration
		}
	}
}
