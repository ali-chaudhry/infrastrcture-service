package main

import "fmt"

func main() {
	var conferenceName = "Go conference"
	const tickets = 50
	fmt.Printf("Welcome to %v booking application. \n", conferenceName)
	fmt.Printf("Only %v tickets left.", tickets)
}