package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50
	var bookings = []string{}
	// or: var bookings [50]string
	// if you add a value to [] - it is array (fixed size)
	// if no value added to [] - it is slice (dynamic in size)

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v of 2022.\n", conferenceName)
	fmt.Printf("There is %v total tickets and there is %v remaining tickets.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets to attend!")


	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		// ask user for their details
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email: ")
		fmt.Scan(&email)
		fmt.Println("How many tickets do you want to purchase: ")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName + " " + lastName )

		fmt.Printf("Thank you %v %v. You have booked %v tickets. Confirmation email sent to %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		firstNames := []string{}
		// _ is used in Go to identify unused variables, since in Go all named variables must be used.
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("The first names of bookings are: %v\n", firstNames)
		
		if remainingTickets == 0 {
			// end program
			
		}

	}
}
