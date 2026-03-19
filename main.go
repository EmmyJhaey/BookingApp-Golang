// package declaration
package main

//importing neccessary packages
import "fmt"

//main function
func main(){
	//constant declaration
	const conferenceTickets int = 50;
	const conferenceName string = "Go Conference";

    //variable declaration
	var firstName string;
	var lastName string;
	var email string;
	var ticketBook int;
	// var remainingTickets int;
	var users = [] string {};

	//Print types of each variables
	/*
		fmt.Println("The types of each variables are: ")
		fmt.Printf("FirstName: %T, LastName: %T, Email: %T, TicketBook: %T, RemainingTickets: %T\n", firstName, lastName, email, ticketBook, remainingTickets)
	*/


	for{
		//Welcome message
		fmt.Printf("Welcome to %v\n", conferenceName)
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
		// remainingTickets = conferenceTickets - ticketBook;
	
		//populating the user array 
		// users[2] = firstName + " " + lastName;
		users = append(users, firstName + " " + lastName)//appending the user input to the users slice 

		//Printing the names of the users who have booked tickets
		fmt.Printf("These are the name of the users who have booked tickets %v\n", users)

		



	
		//booking confirmation message
		// fmt.Printf("Thank you %v %v, for booking %v tickets. A confirmation message has been sent to %v\n", firstName, lastName, ticketBook, email)
		// fmt.Printf("we still have %v tickets available for %v\n", remainingTickets, conferenceName)
	
		//Printing the names of the users who have booked tickets
		// fmt.Printf("The names of the users who have booked tickets are: %v\n", users)
		//Print Slice length
		// fmt.Printf("The length of the users slice is: %v\n", len(users))
	}
}
