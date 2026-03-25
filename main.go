// package declaration
package main

//importing neccessary packages
import (
	"fmt"
	"strconv"
)

//package level variable
const conferenceName string = "Go Conference";
const conferenceTickets int = 50;
var remainingTickets = 50;
var bookings = make([]map[string]string, 0);


//main function
func main(){

	//welcome message into Go Conference Application
	welcomeMessage()

	for{
		//taking user input
		firstName, lastName, email, userTicket := getUserInput()

		//validate user input 
		isValidName, isValidEmail, isValidTicketNumber := validateUserInputs(firstName, lastName, email, int(userTicket))

		//if user inputs invalid value
		if  isValidName && isValidEmail && isValidTicketNumber {

			//process ticket
			processTicket(int(userTicket));

			//get user's data 
			userData := make(map[string]string)

			userData["firstName"] = firstName;
			userData["lastName"] = lastName;
			userData["email"] = email;
			userData["userTicket"] = strconv.FormatInt(int64(userTicket), 10);


			//test our map - by print user info 
			fmt.Printf("User Info %v: \n=================\n", userData)
			
			fmt.Printf("Thank you, %v for booking %v ticket. A mail will be sent to %v\n", userData["firstName"], userData["userTicket"], userData["email"])
			fmt.Printf("%v\n", lines)
			fmt.Printf("We have %v more ticket remaining for the %v\n", remainingTickets, conferenceName)

			//append user data to booking slice
			bookings = append(bookings, userData)

			//print users first name only 
			
			firstNames := getFirstNames()
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
			}
			if !isValidName{
				fmt.Printf("Your First Name or Last should be more than 2 Characters Long\n")
			}
			if !isValidTicketNumber{
				fmt.Printf("Ensure to book within the scope of the available ticket(%v)\n", remainingTickets)
			}
			fmt.Printf("==================================================================\n")

			continue
		}
	}
}












