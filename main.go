package main

import "fmt"

func main() {
	conferenceName := "Go conference"
	const tickets int = 50
	// uinit cannot have a negative value 
	var remainingTickets uint = 50;

	fmt.Printf("Welcome to %v booking application. Total tickets are %v . \n", conferenceName, tickets)

	var userName string
	var userTickets int
	userName = "Ali"
	userTickets = 10
	remainingTickets -= uint(userTickets);

	fmt.Printf("User %v has bought %v tickets. ",userName, userTickets)
	fmt.Printf("Only %v tickets left. \n", remainingTickets)



}