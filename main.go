package main

import "fmt"

func main() {
	var conferenceName = "Go conference"
	const tickets = 50
	fmt.Println("Welcome to ", conferenceName, "booking application")
	fmt.Println("Only ", tickets, " tickets left")
}