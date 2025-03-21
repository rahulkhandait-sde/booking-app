package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = conferenceTickets

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var userName string
	var userTickets int
	//ask user to enter their name
	fmt.Print("Enter your name: ")
	fmt.Scanln(&userName)
	//ask user to enter number of tickets they want to book
	fmt.Printf("Hello %v, how many tickets you want to book?\n", userName)
	fmt.Scanln(&userTickets)
	fmt.Printf("You have booked %v tickets for %v\n", userTickets, conferenceName)

}