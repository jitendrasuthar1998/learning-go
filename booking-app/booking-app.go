package main

import "fmt"

func main() {
	var conferenceName string = "Go conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	// fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	// fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets int
	// ask user for their name
	fmt.Println("Enter your first name")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets")
	fmt.Scan(&userTickets)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confimation email at %v\n.", firstName, lastName, userTickets, email)

	// fmt.Printf("User full name is %v %v and email is %v.\n", firstName, lastName, email)
	// userName = "Jitubhai"
	// userTickets = 2
	// fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)
}
