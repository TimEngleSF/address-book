package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Address struct {
	Address string
	City    string
	State   string
	ZipCode string
}

type Contact struct {
	Name    string
	Email   string
	Phone   string
	Address Address
}

type contacts []Contact

func getContactInfo() Contact {
	c := Contact{}

	scanner = bufio.NewScanner(os.Stdin)

	fmt.Println("Name: ")
	scanner.Scan()
	c.Name = strings.TrimSpace(scanner.Text())

	fmt.Println("Email: ")
	scanner.Scan()
	c.Email = strings.TrimSpace(scanner.Text())

	fmt.Println("Phone: ")
	scanner.Scan()
	c.Phone = strings.TrimSpace(scanner.Text())

	fmt.Println("Street Address:")
	scanner.Scan()
	c.Address.Address = strings.TrimSpace(scanner.Text())

	fmt.Println("City:")
	scanner.Scan()
	c.Address.City = strings.TrimSpace(scanner.Text())

	fmt.Println("State:")
	scanner.Scan()
	c.Address.State = strings.TrimSpace(scanner.Text())

	fmt.Println("Zip Code:")
	scanner.Scan()
	c.Address.ZipCode = strings.TrimSpace(scanner.Text())

	return c
}

func (cs *contacts) add(C *Contact) {
	if C == nil {
		c := getContactInfo()
		*cs = append(*cs, c)
	} else {
		*cs = append(*cs, *C)
	}
	cs.saveToFile()
}

func (cs *contacts) viewContacts() {
	fmt.Printf("\n")
	if len(*cs) == 0 {
		fmt.Println("No contacts have been added")
	} else {
		scanner = bufio.NewScanner(os.Stdin)

		for i, s := range *cs {
			fmt.Printf("%d: %s\n", i+1, s.Name)
		}

		fmt.Println("\nPlease Select a Contact")
		scanner.Scan()
		inputStr := strings.TrimSpace(scanner.Text())
		inputInt, err := strconv.Atoi(inputStr)
		if err != nil {
			fmt.Printf("Error converting input to integer, %s", err)
		}

		(*cs)[inputInt-1].print()
	}
	fmt.Printf("\n")
}

func (cs *contacts) deleteContact() {
	fmt.Printf("\n")
	if len(*cs) == 0 {
		fmt.Println("No contacts have been added")
	} else {
		scanner = bufio.NewScanner(os.Stdin)

		for i, s := range *cs {
			fmt.Printf("%d: %s\n", i+1, s.Name)
		}

		fmt.Println("\nPlease Select a Contact")
		scanner.Scan()
		inputStr := strings.TrimSpace(scanner.Text())
		inputInt, err := strconv.Atoi(inputStr)
		if err != nil {
			fmt.Printf("Error converting input to integer, %s", err)
		}

		*cs = append((*cs)[:inputInt-1], (*cs)[inputInt:]...)
	}
	cs.saveToFile()
	fmt.Printf("\n")
}

func (c Contact) print() {
	fmt.Printf("\n")
	fmt.Println("Name:", c.Name)
	fmt.Println("Email:", c.Email)
	fmt.Println("Phone:", c.Phone)
	fmt.Println("Address:", c.Address.Address)
	fmt.Println("City:", c.Address.City)
	fmt.Println("State:", c.Address.State)
	fmt.Println("Zip Code:", c.Address.ZipCode)
	fmt.Printf("\n")
}

func getContactsFile(cs *contacts) error {
	data, err := os.ReadFile("./contacts.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, cs)
	if err != nil {
		return err
	}
	return nil
}

func (cs *contacts) saveToFile() error {
	data, err := json.Marshal(cs)
	if err != nil {
		return err
	}

	err = os.WriteFile("./contacts.json", data, 0644)
	if err != nil {
		return err
	}

	return nil
}
