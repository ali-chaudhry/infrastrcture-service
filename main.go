package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go conference"
	const tickets int = 50
	// uinit cannot have a negative value
	var remainingTickets uint = 50
	bookings := []string{}
	fmt.Printf("Welcome to %v booking application. Total tickets are %v . \n", conferenceName, tickets)

	var userName string
	var userTickets int
	var purchaseRatio float64

	userName = "Ali"
	userTickets = 10
	remainingTickets -= uint(userTickets)
	purchaseRatio = (float64(remainingTickets) / float64(tickets)) * 100

	fmt.Printf("User %v has bought %v tickets. ", userName, userTickets)
	fmt.Printf("Only %v tickets left. \n", remainingTickets)
	fmt.Printf("Sold %v pecent of the the tcikets \n", purchaseRatio)

	// conditional for loop
	for remainingTickets > 0 && len(bookings) < 50 {

		firstName, lastName, email, userTickets := getInput()
		isValidName, isValidEmail, isValidTicketNumbers := validate(firstName, lastName, email, userTickets, remainingTickets)

		if !isValidName {
			fmt.Printf("FirstName %v or LastName %v is not valid", firstName, lastName)
			continue
		}

		if !isValidEmail {
			fmt.Printf("Email %v is not valid", email)
			continue
		}

		if !isValidTicketNumbers {
			fmt.Printf("Only %v tickets left", remainingTickets)
			continue
		}
		remainingTickets -= userTickets
		bookings = append(bookings, firstName+"  "+lastName)

		greetUser(firstName, lastName, int(userTickets), email, int(remainingTickets), conferenceName)

		firstNames := getFirstNames(bookings)

		fmt.Printf("These are all our bookings: %v \n", firstNames)

		if remainingTickets == 0 {
			// end prgram
			fmt.Println("All tickets are booked")
			break
		}

		city := "London"

		switch city {
		case "London":
			// code for London
		case "New York":
			// code for New York city
		case "Singapor":
			// code for Singapor city
		default:
			// default code
			fmt.Print("No valid city")
		}

	}

}

func getInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Println("Enter your tickets:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func greetUser(firstName string, lastName string, userTickets int, email string, remainingTickets int, conferenceName string) {
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive confirmation email at %v .\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)
}

func getFirstNames(bookings []string) []string {
	// alrenate way of declaring
	firstNames := []string{}
	// for each loop  example
	// _ is used for ignoring index
	// for index, booking
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}

	return firstNames
}

func validate(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {

	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumbers := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumbers
}
