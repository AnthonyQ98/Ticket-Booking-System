package main

import (
	"fmt"
	"strings"
	"booking-system/helper"
)

var conferenceName = "Go Conference"
const conferenceTickets uint = 50
var remainingTickets uint = 50
var bookings = []string{}

func main() {
	// call function
	greetUser()

	for {
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketAmount := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketAmount {
			bookTicket(userTickets, firstName, lastName, email)

			// print first names in function
			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)
			
			noTicketsRemaining := remainingTickets == 0
			// or "if remainingTickets == 0 {}" would also work below.
			if noTicketsRemaining {
				// end program
				fmt.Println("SORRY :(! Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("ERROR: Name is invalid. Please try again.")
			}
			if !isValidEmail {
				fmt.Println("ERROR: Email is invalid. Please try again.")
			}
			if !isValidTicketAmount {
				fmt.Println("ERROR: Ticket amount is invalid. please try again.")
			}
		}
		
	}

	/*
	city := "London"

	switch city {
		case "New York", "Chicago": // multiple cases.
			// execute code
		case "Singapore":
			// execute diff code.
		case "Dublin", "Cork": // can add multiple cases for situations where same code should be executed.
			// execute diff code 2.
		default:
			// when none of the above cases are true, execute default block.
	}
	*/
}


func greetUser() {
	fmt.Printf("Welcome to the %v booking application.\n", conferenceName)
	fmt.Printf("There is %v total tickets and there is %v remaining tickets.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets to attend!")
}


func getFirstNames() []string {
	firstNames := []string{}
	// _ is used in Go to identify unused variables, since in Go all named variables must be used.
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	//fmt.Printf("The first names of bookings are: %v\n", firstNames)
	return firstNames
}

func getUserInput() (string, string, string, uint) {
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
	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) (uint, []string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName + " " + lastName )

	fmt.Printf("Thank you %v %v. You have booked %v tickets. Confirmation email sent to %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
	return remainingTickets, bookings

}