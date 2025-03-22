package main

import (
	"strings"
)

func  ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	isValidDetails := isValidName && isValidEmail && isValidTicketNumber
	return isValidDetails, isValidName, isValidEmail, isValidTicketNumber
}