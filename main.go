package main

import (
	"fmt"
	"strings"
	"time"
)

/*
TICKET RESERVATION SYSTEM:
The program continuously asks user inputs to determine the event that a user wants to attend or
exit the service followed by the user details as in name, age, gender, and email along with the
number of tickets to be bought. The program validates the user inputs and prints the tickets as
an output while maintaining the reservation list which is printed at the end.

The program uses go routines to separate the thread of printing tickets to the user and struct
array to store the reservation list with user details.

The project is used for learning the basics of GOLANG programming language.
*/

type list struct {
	tickets int
	name    string
	age     int
	gender  string
	email   string
	event   string
}

var reservation = make([]list, 0, 50)

func main() {

	const totalTickets int = 50
	var event_1_seats_left int = 50
	var event_2_seats_left int = 50
	var event_3_seats_left int = 50
	var event_4_seats_left int = 50
	var ticketsBought int = 0
	var event string
	startService()
	for {
		var choice = welcome()
		if choice == 0 {
			endService()
			print_reservation_list(reservation)
			return
		} else if choice == 1 {
			event_1_seats_left, ticketsBought, event = choices(choice, totalTickets, event_1_seats_left)
		} else if choice == 2 {
			event_2_seats_left, ticketsBought, event = choices(choice, totalTickets, event_2_seats_left)
		} else if choice == 3 {
			event_3_seats_left, ticketsBought, event = choices(choice, totalTickets, event_3_seats_left)
		} else if choice == 4 {
			event_4_seats_left, ticketsBought, event = choices(choice, totalTickets, event_4_seats_left)
		}
		if ticketsBought == 0 {
			continue
		}
		fmt.Println("Please enter your details:")
		var name, age, gender, email = confirmDetails(inputs())
		fmt.Println("Here are the booking details for the ticket")
		printDetail(name, email, ticketsBought)
		printTickets(name, age, gender, email, ticketsBought, event)
		fmt.Println("\nWould you like to book tickets for another event?")
	}
}

func welcome() int {
	var totalEvents int = 4
	// fmt.Println()
	// fmt.Println("**************************************************************************************************************")
	// fmt.Println("                                  WELCOME TO THE TICKET BOOKING SERVICE")
	// fmt.Println("Please select the event you want to attend:")
	fmt.Println("1 ---> Python Conference")
	fmt.Println("2 ---> Golang Conference")
	fmt.Println("3 ---> Java Conference")
	fmt.Println("4 ---> C++ Conference")
	fmt.Println("0 ---> Exit")
	fmt.Println()

	var choice int
	fmt.Print("Your choice : ")
	fmt.Scan(&choice)

	for {
		if choice < 0 || choice > totalEvents {
			fmt.Println("Invalid Choice. Please Try Again!!!!")
			fmt.Print("Your choice : ")
			fmt.Scan(&choice)
		} else {
			break
		}
	}
	return choice
}

func inputs() (string, int, string, string) {
	var name string
	var age int
	var gender string
	var email string

	for {
		fmt.Print("\nEnter your name : ")
		fmt.Scan(&name)
		if len(name) < 2 || len(name) > 100 {
			fmt.Print("Invalid name. Please Try Again!\n\n")
		} else {
			break
		}
	}

	for {
		fmt.Print("Enter your age : ")
		fmt.Scan(&age)
		if age < 0 || age > 150 {
			fmt.Print("Invalid age. Please Try Again!\n\n")
		} else {
			break
		}
	}

	for {
		fmt.Print("Enter your gender [M/F]: ")
		fmt.Scan(&gender)
		if gender != "m" && gender != "f" && gender != "M" && gender != "F" {
			fmt.Print("Invalid gender. Please Try Again!\n\n")
		} else {
			if gender == "m" || gender == "M" {
				gender = "MALE"
			} else {
				gender = "FEMALE"
			}
			break
		}
	}

	for {
		fmt.Print("Enter your email : ")
		fmt.Scan(&email)
		if !strings.Contains(email, "@") || len(email) < 6 || !strings.Contains(email, ".com") {
			fmt.Print("Invalid email. Please Try Again!\n\n")
		} else {
			break
		}
	}

	return name, age, gender, email
}

func confirmDetails(name string, age int, gender string, email string) (string, int, string, string) {
	fmt.Println("\nAre your details below correct?")
	fmt.Printf("Name   : %v\n", name)
	fmt.Printf("Age    : %v\n", age)
	fmt.Printf("Gender : %v\n", gender)
	fmt.Printf("Email  : %v \n\n", email)
	fmt.Print("CONFIRM [Y/N] :")

	var confirm string
	var name1 string
	var age1 int
	var gender1 string
	var email1 string

	fmt.Scan(&confirm)
	if confirm == "n" || confirm == "N" {
		fmt.Println("Sorry to hear that! Please re-enter your details.")
		name1, age1, gender1, email1 = inputs()
		name, age, gender, email = confirmDetails(name1, age1, gender1, email1)
	} else if confirm != "Y" && confirm != "y" && confirm != "N" && confirm != "n" {
		fmt.Println("Invalid entry. Please Try Again!")
		confirmDetails(name, age, gender, email)
	}
	return name, age, gender, email
}

func endService() {
	fmt.Println("\nThank you for using our service. We hope to see you again!")
	fmt.Println("**************************************************************************************************************")
	fmt.Println()
}

func startService() {
	fmt.Println()
	fmt.Println("**************************************************************************************************************")
	fmt.Println("                                  WELCOME TO THE TICKET BOOKING SERVICE")
	fmt.Println("Please select the event you want to attend:")
}

func confirmTickets(event string, totalTickets int, ticketsLeft int) int {
	fmt.Print("\n", event, "\n")
	fmt.Printf("\nThere are %v seats for the event and %v tickets are left\n", totalTickets, ticketsLeft)
	fmt.Println("Would you like to buy tickets for this event? [Y/N]")
	fmt.Print("Confirmation = ")
	var confirm string
	var tickets int
	fmt.Scan(&confirm)
	for {
		if confirm == "N" || confirm == "n" {
			return 0
		} else if confirm == "Y" || confirm == "y" {
			fmt.Println("Enter the number of tickets you want to buy?")
			fmt.Print("Tickets = ")
			fmt.Scan(&tickets)
			for {
				if tickets > ticketsLeft || tickets < 1 {
					fmt.Println("Invalid number of tickets. Please Try Again!")
					fmt.Print("Tickets = ")
					fmt.Scan(&tickets)
				} else {
					fmt.Printf("\nConfirm %v tickets for %v\n\n", tickets, event)
					break
				}
			}
			break
		} else {
			fmt.Println("Invalid Input. Please Try Again!")
			fmt.Println("Would you like to buy tickets for this event? [Y/N]")
			fmt.Print("Confirmation = ")
			fmt.Scan(&confirm)
		}
	}
	return tickets
}

func choices(choice int, totalTickets int, ticketsLeft int) (int, int, string) {
	var ticketsBought int
	var event string
	switch choice {
	case 1:
		event = "Event 1 : Python Conference"
		ticketsBought = confirmTickets(event, totalTickets, ticketsLeft)
		if ticketsBought == 0 {
			fmt.Println("\nSorry to hear you don't want to attend the event. Please choose an option!")
			break
		} else {
			ticketsLeft -= ticketsBought
		}

	case 2:
		fmt.Println()
		event = "Event 2 : Golang Conference"
		ticketsBought = confirmTickets(event, totalTickets, ticketsLeft)
		if ticketsBought == 0 {
			fmt.Println("\nSorry to hear you don't want to attend the event. Please choose an option!")
			break
		} else {
			ticketsLeft -= ticketsBought
		}

	case 3:
		event = "Event 3 : Java Conference"
		ticketsBought = confirmTickets(event, totalTickets, ticketsLeft)
		if ticketsBought == 0 {
			fmt.Println("\nSorry to hear you don't want to attend the event. Please choose an option!")
			break
		} else {
			ticketsLeft -= ticketsBought
		}

	case 4:
		event = "Event 4 : C++ Conference"
		ticketsBought = confirmTickets(event, totalTickets, ticketsLeft)
		if ticketsBought == 0 {
			fmt.Println("\nSorry to hear you don't want to attend the event. Please choose an option!")
			break
		} else {
			ticketsLeft -= ticketsBought
		}
	}
	return ticketsLeft, ticketsBought, event
}

func printDetail(name string, email string, ticketsBought int) {
	fmt.Printf("\t%v tickets confirmed for Mr/Mrs %v\n", ticketsBought, name)
	fmt.Printf("Confirmation email was sent to \n\t\t%v\n", email)
	fmt.Println("A copy of your tickets will be generated soon!")
	fmt.Printf("ENJOY YOUR EVENT!\n\n")
}

func printTickets(name string, age int, gender string, email string, ticketsBought int, event string) {
	// time.Sleep(10000)
	var list1 list
	list1.name = name
	list1.age = age
	list1.gender = gender
	list1.email = email
	list1.event = event
	list1.tickets = ticketsBought
	reservation = append(reservation, list1)
	var current_time = time.Now()
	fmt.Println("##################################################################################################################")
	fmt.Printf("##                                            %v\n", event)
	fmt.Printf("##    NAME   : %v\n", name)
	fmt.Printf("##    AGE    : %v             EMAIL : %v\n", age, email)
	fmt.Printf("##    GENDER : %v           TICKET COUNTS : %v\n", gender, ticketsBought)
	fmt.Println("##")
	fmt.Printf("##                                                  Time : %v\n", current_time)
	fmt.Println("##################################################################################################################")
}

func print_reservation_list(reservation []list) {
	fmt.Println("\nBelow is the reservation list for the events: ")
	for i := range reservation {
		fmt.Println(reservation[i])
	}
	fmt.Print("\n\n")
}
