package main

import (
	"fmt"
	"sync"
	"time"
	"yet-another-booking-app/helper"
)

// CREATE A VARIABLE AS LOCAL CAS POSSIBLE
// cant use syntactic sugar on package level vars
// var & consts
var bookingAppName = "The greatest Go booking app ever" //syntactic sugar
const totalTickets = 50

var remainingTickets uint = 50

// array & slices
var bookings = make([]UserData, 0)

type UserData struct {
	firstName string
	lastName  string
	email     string
	tickets   uint
}

var wg = sync.WaitGroup{}

func main() {

	fmt.Printf("bookingAppName is %T, totalTickets is %T, remainingTickets is %T\n", bookingAppName, totalTickets, remainingTickets)

	greetUser()

	//for: only loop
	//for {

	firstName, lastName, email, userTickets := getUserInput()

	isValidName, isValidEmail, isValidAmount := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidAmount && isValidName && isValidEmail {

		bookTicket(userTickets, firstName, lastName, email)

		wg.Add(1)
		go sendTicket(firstName, lastName, email, userTickets)

		firstNames := printFirstNames()
		fmt.Printf("All bookings: %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Printf("No more tickets available.\n")
			//break
		}
	} else {
		if !isValidName {
			fmt.Printf("Your name is too short.\n")
		}
		if !isValidEmail {
			fmt.Printf("Your email is invalid.\n")

		}
		if !isValidAmount {
			fmt.Printf("Invalid ticket number.\n")

		}
		//continue //next iteration
	}
	wg.Wait()
}

//}

func greetUser() {
	//println & printf
	fmt.Printf("Welcome to %v, enjoy your stay\n", bookingAppName)
	fmt.Println("You are about to buy tickets")
	fmt.Printf("Total tickets: %v, Available tickets: %v\n", remainingTickets, totalTickets)
}

func printFirstNames() []string {
	firstNamesSlice := []string{}
	for _, element := range bookings { //range iterates // _ : blank identifier //split by space
		firstNamesSlice = append(firstNamesSlice, element.firstName)
	}
	return firstNamesSlice
}

func getUserInput() (string, string, string, uint) {
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

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets -= userTickets

	var userData = UserData{
		firstName: firstName,
		lastName:  lastName,
		email:     email,
		tickets:   userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("Bookings: %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive your order details at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remainig for %v.\n", remainingTickets, bookingAppName)
}

func sendTicket(firstName string, lastName string, email string, userTickets uint) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v\n", firstName, lastName, userTickets)
	fmt.Println("--------------------------------")
	fmt.Printf("Sending ticket: \n%v\n to %v\n", ticket, email)
	fmt.Println("--------------------------------")

	wg.Done()
}
