package main

import (
	"fmt"
	"strings"
)

func main() {
	confrenceName := "Go Confrence"
	const confrenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Printf("Welcome to %v booking appplication\n", confrenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", confrenceTickets, remainingTickets)
	fmt.Println("Get yout tickets here to attend")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		// ask user for their name
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		if userTickets <= remainingTickets {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thanks you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, confrenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("The first names of bookings are: %v\n", bookings)

			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is boked our. Come back next year.")
				break
			}
		} else
		
		fmt.Printf("We only have %v tickets remaining, so you can't %v tickets\n", remainingTickets, userTickets)
		continue

	}

}
