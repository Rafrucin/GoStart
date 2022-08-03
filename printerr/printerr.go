package printerr

import (
	"gostart/models"
)

func PrintTickets(tickets []models.UserData ) (string, string) {
	if len(tickets) < 2 {
		return "", "not too many - not worth to start"
	}
	return "Printed", ""
}

