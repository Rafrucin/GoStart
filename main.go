package main

import (
	"fmt"
	"gostart/printerr"
	"reflect"
	"strings"
	"gostart/models"
)

const confTickets = 50

var remainingTickets = 50
var firstName string
var lastName string
var email string
var userTickets uint
var userList = make([]models.UserData, 0)

func main() {

	var confName = "RRCONF"

	var bookings []string

	for remainingTickets > 0 {
		fmt.Printf("Welcome to our conf %v \n", confName)
		fmt.Println("We have a total of", confTickets, "and", remainingTickets, "still remaining")

		fmt.Println("What's your first name?")
		fmt.Scanln(&firstName)
		fmt.Println("What's your last name?")
		fmt.Scanln(&lastName)

		fmt.Println("What's your email?")
		fmt.Scanln(&email)

		checks := checks()

		if checks != "" {
			fmt.Println(checks)
			continue
		}

		fmt.Println("Howmany tickets do you need?")
		fmt.Scanln(&userTickets)
		fmt.Println(userTickets)
		if userTickets == 0 {
			fmt.Println("Ticket number not valid")
			continue
		}

		if !validateTicketsNumber(int(userTickets)) {
			println("Not enough tickets remaining.")
			continue
		}
		remainingTickets -= int(userTickets)

		bookings = append(bookings, firstName+" "+lastName)

		fmt.Println("Welcome", firstName, lastName, "to", confName, "you've booked", userTickets, "tickets. Remaining tickets:",
			remainingTickets, reflect.TypeOf(bookings), len(bookings), ". Enjoy", bookings[0])

		fmt.Printf("Slice items: %v \n", bookings)

		for index, booking := range bookings {
			fmt.Printf("slice element %v is %v \n", index, booking)
		}

		var firstNames []string
		for _, booking := range bookings {
			firstNameOnly := strings.Fields(booking)
			firstNames = append(firstNames, firstNameOnly[0])
		}

		fmt.Printf("%v booked tickets \n\n", firstNames)

		var userOrder = models.UserData{
			FirstName: firstName,
			LastName: lastName,
			Email: email,
			NumberOfTickets: userTickets,
		}

		userList = append(userList, userOrder)
		output, err := printerr.PrintTickets(userList)

		if output == "" {
			println(err)
		} else {
			println(output)
		}


	}
}

func checks() string {

	if len(firstName) < 2 || len(lastName) < 2 {
		return "name not long enough"
	}

	if !strings.Contains(email, ".") || !strings.Contains(email, "@") {
		return "email not valid"
	}

	return ""
}
