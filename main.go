package main

import (
	"fmt"
	"strings"
)

// Constants
const ConferenceName = "Go Conference"
const NumberOfTickets uint = 50

// Structs
type users struct {
	FirstName       string
	LastName        string
	Email           string
	NumberOfTickets uint
	Country         string
}

type Country struct {
	Name             string
	Place            string
	Date             string
	RemainingTickets uint
}

// Data
var Countries = []Country{
	{"Ethiopia", "Gondar, Piazza", "AUG 11", 50},
	{"England", "London Hall, UK", "AUG 15", 50},
	{"France", "Paris Central", "AUG 18", 50},
	{"Germany", "Berlin Arena", "AUG 20", 50},
	{"Russia", "Moscow Red Stage", "AUG 23", 50},
	{"USA", "New York Conf Center", "AUG 25", 50},
}

var Booking = []users{} // Booked users

// Function to greet the user
func GreetUsers() {
	fmt.Printf("Welcome to %v booking application!\n\n", ConferenceName)
}

// Function to validate inputs
func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

// Booking logic
func BookTicket(user users, country *Country) {
	country.RemainingTickets -= user.NumberOfTickets
	Booking = append(Booking, user)
	fmt.Printf("\n✅ Thank you, %v %v, for booking %v ticket(s).\n", user.FirstName, user.LastName, user.NumberOfTickets)
	fmt.Printf("📩 A confirmation email will be sent to: %v\n", user.Email)
	fmt.Printf("🎟️ Tickets remaining for %v: %v\n", country.Name, country.RemainingTickets)
	fmt.Println("--------------------------------------------------")
}

// Main logic
func main() {
	GreetUsers()

	for {
		// Show available countries
		fmt.Println("Available Countries:")
		for i, country := range Countries {
			fmt.Printf("%v. %v\n", i+1, country.Name)
		}

		// Select country
		var CountryChoice int
		fmt.Print("\nEnter your choice number: ")
		fmt.Scan(&CountryChoice)

		if CountryChoice < 1 || CountryChoice > len(Countries) {
			fmt.Println("❌ Invalid country selection.")
			continue
		}

		// Get pointer to selected country 
		selectedCountry := &Countries[CountryChoice-1]

		// Show country details
		fmt.Printf("\n==== %v Go Conference ====\n", selectedCountry.Name)
		fmt.Printf("📍 Location: %v\n", selectedCountry.Place)
		fmt.Printf("📅 Date: %v\n", selectedCountry.Date)
		fmt.Printf("🎟️ Tickets Remaining: %v\n\n", selectedCountry.RemainingTickets)

		// Get user input 
		var firstName, lastName, email string
		var userTickets uint

		fmt.Print("Enter your first name: ")
		fmt.Scan(&firstName)
		fmt.Print("Enter your last name: ")
		fmt.Scan(&lastName)
		fmt.Print("Enter your email: ")
		fmt.Scan(&email)
		fmt.Print("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		// Validate input
		isValidName, isValidEmail, isValidTicketNumber := ValidateUserInput(firstName, lastName, email, userTickets, selectedCountry.RemainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			user := users{
				FirstName:       firstName,
				LastName:        lastName,
				Email:           email,
				NumberOfTickets: userTickets,
				Country:         selectedCountry.Name,
			}
			BookTicket(user, selectedCountry)
		} else {
			if !isValidName {
				fmt.Println("❌ First or last name must be at least 2 characters.")
			}
			if !isValidEmail {
				fmt.Println("❌ Invalid email format.")
			}
			if !isValidTicketNumber {
				fmt.Printf("❌ You requested %v tickets, but only %v are available.\n", userTickets, selectedCountry.RemainingTickets)
			}
			continue
		}

		// Ask to book again
		var Continue int
		fmt.Println("\nDo you want to book another ticket?\n1. Yes\n2. No")
		fmt.Scan(&Continue)

		if Continue == 2 {
			fmt.Println("\n👋 Thank you for using the Go Conference Booking App!")
			break
		} else if Continue != 1 {
			fmt.Println("❌ Invalid choice. Exiting.")
			break
		}
	}
}
