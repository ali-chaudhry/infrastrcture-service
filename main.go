package main

import "fmt"

func main() {
	var conferenceName string = "Go conference"
	const tickets int = 50
	fmt.Printf("Welcome to %v booking application. \n", conferenceName)
	fmt.Printf("Only %v tickets left. \n", tickets)

	var userName string
	var userTickets int
	userName = "Ali"
	userTickets = 10

	fmt.Printf("User %v has bought %v tickets. ",userName, userTickets)



}