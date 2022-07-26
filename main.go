package main

import "fmt"

func main() {
	var confName = "RRCONF"
	const confTickets = 50
	remainingTickets := 50
	var userName string
	var userTickets uint

	fmt.Printf("Welcome to our conf %v \n", confName)
	fmt.Println("We have a total of", confTickets, "and", remainingTickets, "still remaining")

	fmt.Println("What's your name?")
	fmt.Scanln(&userName)
	fmt.Println("Howmany tickets do you need?")
	fmt.Scanln(&userTickets)
	fmt.Println("Welcome", userName, "to",confName,"you've booked", userTickets, "tickets")
}
