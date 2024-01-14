package helper

import "strings"

//scope: package level, vars anda funcs can be accessed in other files within the same package

//TO EXPORT JUST CAPITALIZE FIRST LETTER
func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidAmount := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidAmount
}
