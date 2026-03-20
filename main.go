// package declaration
package main

//importing neccessary packages
import (
	"fmt"
	"strings"
)

//main function
func main(){
	//Constant Declaration
	const conferenceName string = "Go Conference";

    //Variable declaration
	conferenceTickets := 50;
	var firstName string;
	var lastName string;
	var email string;
	var ticketBook int;
	var remainingTickets int;
	bookings := [] string {};



	for{
		//Welcome message
		fmt.Printf("Welcome to %v\n", conferenceName)
		fmt.Printf("==================================================================\n")
		fmt.Printf("We have total of %v tickets, and we are here at your service...\n", conferenceTickets)
	
		//taking user input
		//user input for first name
		fmt.Printf("What is your First Name \n")
		fmt.Scan(&firstName)
	
		//user input for last name
		fmt.Printf("What is your last Name \n")
		fmt.Scan(&lastName)
	
		//user input for first name
		fmt.Printf("What is your email address \n")
		fmt.Scan(&email)
	
		//user input for ticket booking
		fmt.Printf("How many ticktes would you like to book? \n")
		fmt.Scan(&ticketBook)
	
		//calculating remaining tickets
		remainingTickets = conferenceTickets - ticketBook;

		//Final Output
		fmt.Printf("Thank you, %v for booking %v ticket. A mail will be sent to %v\n", firstName, ticketBook, email)
		fmt.Printf("We have %v more ticket remaining for the %v\n", remainingTickets, conferenceName)


		//list of bookings 
		bookings = append(bookings, firstName + " " + lastName);

		firstNames := [] string {};

		for _, booking := range bookings{
			names := strings.Fields(booking)
			//print out users first name only
			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("The first name of bookings: %v\n", firstNames)

		conferenceTickets -= ticketBook;
		fmt.Printf("==================================================================\n")
	}
}
