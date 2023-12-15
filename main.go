package main

import	"fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50


	fmt.Printf("data type of conferenceName %T, conferenceTickets %T, remainingTickets %T\n", conferenceName, conferenceTickets, remainingTickets)
	
	fmt.Println("Welcome to our GO conference booking application.")
	fmt.Println("Get your tickets here to attend.")
	fmt.Printf("We have a total of %v and we have %v remaining\n", conferenceTickets, remainingTickets)
	
	fmt.Println((conferenceName))


	var firstName string
	var lastName string
	var email string
	var userTickets int
 //ask user for the name and ticket
 	fmt.Println("\nPlease enter your first Name: ")
	fmt.Scan(&firstName)

	fmt.Println("\nPlease enter your last Name: ")
	fmt.Scan(&lastName)

	fmt.Println("\nPlease enter your email: ")
	fmt.Scan(&email)


	userTickets = 12
	fmt.Printf("A new user with the following details %v %v and email %v wants to book %v tickets.\n", firstName, lastName, email, userTickets)

	fmt.Printf("Thank you %v for booking with us. You order has been received and details will be emailed to you at %v.\n", firstName,email)

}

