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
	const conferenceTickets int = 50;

    //Variable declaration
	var firstName string;
	var lastName string;
	var email string;
	var ticketBook int;

	remainingTickets:= 50;
	bookings := [] string {};



	for{
		//Welcome message
		fmt.Printf("==================================================================\n")
		fmt.Printf("Welcome to %v Application\n", conferenceName)
		fmt.Printf("==================================================================\n")
		fmt.Printf("We have total of %v tickets, with %v remaining\n",conferenceTickets, remainingTickets)
	
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

		//validating user input
		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := int(ticketBook) > 0 && int(ticketBook) <= remainingTickets

		// fmt.Printf("isNameValid: %v, isEmailValid: %v, isTicketNumberValid: %v\n", isValidName, isValidEmail, isValidTicketNumber)
		// break

		//if user inputs invalid value
		if  isValidName && isValidEmail && isValidTicketNumber{

			//calculating remaining tickets
			remainingTickets = remainingTickets - int(ticketBook);

			//Final Output
			fmt.Printf("Thank you, %v for booking %v ticket. A mail will be sent to %v\n", firstName, ticketBook, email)
			fmt.Printf("We have %v more ticket remaining for the %v\n", remainingTickets, conferenceName)

			//list of bookings 
			bookings = append(bookings, firstName + " " + lastName);
			//security layer for bookings first name only 
			firstNames := [] string {};
			for _, booking := range bookings{
				names := strings.Fields(booking)
				//print out users first name only
				firstNames = append(firstNames, names[0])
			}
			//print out bookings first name only 
			fmt.Printf("The first name of bookings: %v\n", firstNames)

			if remainingTickets == 0 {
			//end the application if this condition is true
			fmt.Printf("You have reach the end of our bookings\n") 
			break
		}
						
		} else {
			//prompt user the available ticket 
			if !isValidEmail{
				fmt.Printf("You Email is Invalid !!!!\n")
			}else if !isValidName{
				fmt.Printf("You Name is Invalid !!!!\n")
			}else{
				fmt.Printf("Your Ticket Number is out of bound !!!\n")
			}
		}

		
		fmt.Printf("==================================================================\n")
	}
}
