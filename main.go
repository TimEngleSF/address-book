package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var scanner *bufio.Scanner

func main() {
	glen := Contact{Name: "Glen", Email: "glen@email.com", Phone: "5105551234", Address: Address{Address: "1444 Rich St", City: "East Bay", State: "Ca", ZipCode: "99999"}}

	cont := contacts{}
	cont.add(&glen)
	scanner = bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Address Book Application")
		fmt.Println("1: View Contacts")
		fmt.Println("2: Add Contact")
		fmt.Println("3: Delete Contact")
		fmt.Println("4: Exit")

		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		switch input {
		case "1":
			cont.viewContacts()
		case "2":
			cont.add(nil)
		case "3":
			cont.deleteContact()
		case "4":
			fmt.Println("Exiting the application")
			return
		default:
			fmt.Println("Invalid option, please try again.")
		}
	}

}
