// package declaration
package main

//importing neccessary packages
import (
	"fmt"
	"strings"
)

//package level variable
const conferenceName string = "Go Conference";
const conferenceTickets int = 50;
var remainingTickets = 50;
var bookings = [] string {};


//main function
func main(){

	//welcome message into Go Conference Application
	welcomeMessage()

	for{
		//taking user input
		firstName, lastName, email, ticketBook := getUserInput()

		//validate user input 
		isValidName, isValidEmail, isValidTicketNumber := validateUserInputs(firstName, lastName, email, ticketBook)

		//if user inputs invalid value
		if  isValidName && isValidEmail && isValidTicketNumber {

			//process ticket
			processTicket(ticketBook, firstName, lastName);
			fmt.Printf("Thank you, %v for booking %v ticket. A mail will be sent to %v\n", firstName, ticketBook, email)
			fmt.Printf("We have %v more ticket remaining for the %v\n", remainingTickets, conferenceName)

			//Print First Names 
			firstNames := getFirstNames(bookings)
			fmt.Printf("The first name of bookings: %v\n", firstNames)
			
			//end the application if this condition is true
			if remainingTickets == 0 {
			fmt.Printf("You have reach the end of our bookings\n") 
			break
		}
						
		} else {
			//prompt user the available ticket 
			if !isValidEmail{
				fmt.Printf("Your Email should have '@' symbol\n")
			}else if !isValidName{
				fmt.Printf("Your First Name or Last should be more than 2 Characters Long\n")
			}else{
				fmt.Printf("Ensure to book within the scope of the available ticket(%v)\n", remainingTickets)
			}
		}
		fmt.Printf("==================================================================\n")
	}
}


func welcomeMessage (){
	fmt.Printf("==================================================================\n")
	fmt.Printf("Welcome to %v Application\n", conferenceName)
	fmt.Printf("==================================================================\n")
	fmt.Printf("We have total of %v tickets, with %v remaining\n",conferenceTickets, remainingTickets)
}


func getUserInput() (string, string, string, int){
	//Variable declaration
	var firstName string;
	var lastName string;
	var email string;
	var ticketBook int;

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

	return firstName, lastName, email, ticketBook
}


func validateUserInputs(firstName string, lastName string, email string, ticketBook int)(bool, bool, bool){
	//validating user input
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := int(ticketBook) > 0 && int(ticketBook) <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}


func processTicket(ticketBook int, firstName string, lastName string)(int, [] string){
	remainingTickets = remainingTickets - int(ticketBook);
	bookings = append(bookings, firstName + " " + lastName);

	return remainingTickets, bookings
}



func getFirstNames(bookings [] string) [] string{
	//security layer for bookings first name only 
	firstNames := [] string {};
	for _, booking := range bookings{
		names := strings.Fields(booking)
		//print out users first name only
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}




