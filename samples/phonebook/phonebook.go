package main

import "fmt"

type Contacts map[string]string

func (c Contacts) store() {
	var name, phonenumber string
	fmt.Print("Enter contact: ")
	fmt.Scan(&name, &phonenumber)
	c[name] = phonenumber
	fmt.Println("Contact saved")
}

func (c Contacts) list() {
	for key, value := range c {
		fmt.Println(key, value)
	}
}

func (c Contacts) find() {
	var name string
	fmt.Print("Enter name: ")
	fmt.Scan(&name)

	if value, found := c[name]; found {
		fmt.Printf("%s has number: %s\n", name, value)
	} else {
		fmt.Printf("%s is not found in the phonebook\n", name)
		fmt.Println("Do you want to add it? y/n")
		var answer string
		if fmt.Scan(&answer); answer == "y" {
			c.store()
		}
	}
}

func main() {
	var command string
	contacts := make(Contacts)
	fmt.Println("Welcome to your phonebook.")

	for {
		fmt.Print("Command >: ")
		fmt.Scan(&command)

		if command == "quit" || command == "exit" {
			break
		} else if command == "help" {
			fmt.Print(`- store: enter a new contact
- list: list all contacts in phonebook
- find: find a contact in phonebook
- quit/exit: exit phonebook
`)
		} else if command == "store" {
			contacts.store()
		} else if command == "list" {
			contacts.list()
		} else if command == "find" {
		} else {
			fmt.Println("Unregconized command. Type `help` for the list of commands")
		}
	}
	fmt.Println("Bye bye")
}
