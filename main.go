package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50


	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v of 2022.\n", conferenceName)
	fmt.Printf("There is %v total tickets and there is %v remaining tickets.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets to attend!")


	var firstName string
	var lastName string
	var email string
	var userTickets int
	// ask user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email: ")
	fmt.Scan(&email)
	fmt.Println("How many tickets do you want to purchase: ")
	fmt.Scan(&userTickets)

	fmt.Printf("%v %v booked %v tickets. Confirmation email sent to %v\n", firstName, lastName, userTickets, email)





}
