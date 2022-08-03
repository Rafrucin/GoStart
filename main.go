package main

import (
	"fmt"
	"gostart/models"
	"gostart/printerr"
	"reflect"
	"strings"
	"sync"
	"time"
)

var Wg = sync.WaitGroup{}

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

		Wg.Add(3)
		go FireDontForget("smt")
		go AwaitMe("solo")

		awaitedValue := AwaitMe("value")

		println("Await:",awaitedValue)

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

		for _, userItem := range userList {
			println(userItem.Email)
		}

		

		break
	}
	Wg.Wait()
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

func AwaitMe (input string) string{
	println("wait started", input)
	time.Sleep(2*time.Second)
	println("wait done", input)
	Wg.Done()
	return "Thanks for waiting " + input
}

func FireDontForget (input string) string {
	println("~~~~~~~~~~~~~", input)
	time.Sleep(4*time.Second)
	println("fire done")
	Wg.Done()
	return "back to main"
}
