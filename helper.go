package main

import (
	"fmt"
	"strings"
)

var lines = "=================================================================="


func welcomeMessage (){
	fmt.Printf("%v\n", lines)
	fmt.Printf("Welcome to %v Application\n", conferenceName)
	fmt.Printf("%v\n", lines)
	fmt.Printf("We have total of %v tickets, with %v remaining\n",conferenceTickets, remainingTickets)
}


func getUserInput() (string, string, string, uint){
	//Variable declaration
	var firstName string;
	var lastName string;
	var email string;
	var userTicket uint;

	fmt.Printf("What is your First Name: ")
	fmt.Scanln(&firstName)

	fmt.Printf("What is your last Name: ")
	fmt.Scanln(&lastName)

	fmt.Printf("What is your email address: ")
	fmt.Scanln(&email)
	
	fmt.Printf("How many ticktes would you like to book: ")
	fmt.Scanln(&userTicket)

	return firstName, lastName, email, userTicket
}


func processTicket(ticketBook int)(int){
	remainingTickets = remainingTickets - int(ticketBook);

	return remainingTickets
}



func validateUserInputs(firstName string, lastName string, email string, ticketBook int)(bool, bool, bool){
	//validating user input
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := int(ticketBook) > 0 && int(ticketBook) <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}



func getFirstNames() []string{
	//security layer for bookings first name only 
	firstNames := [] string {}
	for _, booking := range bookings{
		firstNames = append(firstNames, booking["firstName"])
	}
	return firstNames
}
