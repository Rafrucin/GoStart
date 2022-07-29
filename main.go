package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {

	var confName = "RRCONF"
	const confTickets = 50
	remainingTickets := 50
	var firstName string
	var lastName string
	var userTickets uint
	var bookings []string
	for {
			fmt.Printf("Welcome to our conf %v \n", confName)
			fmt.Println("We have a total of", confTickets, "and", remainingTickets, "still remaining")

			fmt.Println("What's your first name?")
			fmt.Scanln(&firstName)
			fmt.Println("What's your last name?")
			fmt.Scanln(&lastName)
			fmt.Println("Howmany tickets do you need?")
			fmt.Scanln(&userTickets)
			remainingTickets -= int(userTickets)

			bookings = append(bookings, firstName + " " + lastName)

			fmt.Println("Welcome", firstName, lastName, "to",confName,"you've booked", userTickets, "tickets. Remaining tickets:",
							 remainingTickets,reflect.TypeOf(bookings), len(bookings), ". Enjoy", bookings[0])

			fmt.Printf("Slice items: %v \n", bookings)

			for index, booking := range bookings{
				firstNameOnly := strings.Fields(booking)
				fmt.Printf("slice element %v is %v \n", index, firstNameOnly[0])
			}
		}

}
