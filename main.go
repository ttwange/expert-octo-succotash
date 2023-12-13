package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50


	fmt.Printf("data type of conferenceName %T, conferenceTickets %T, remainingTickets %T\n", conferenceName, conferenceTickets, remainingTickets)
	
	fmt.Println("Welcome to our GO conference booking application.")
	fmt.Println("Get your tickets here to attend.")
	fmt.Printf("We have a total of %v and we have %v remaining\n", conferenceTickets, remainingTickets)
	
	fmt.Println((conferenceName))
}

