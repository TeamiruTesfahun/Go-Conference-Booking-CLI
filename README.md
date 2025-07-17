# Go Conference Booking CLI

This is a command-line interface (CLI) application built with Go that lets users book tickets for the "Go Conference" held in various countries. It's a straightforward, interactive program that showcases fundamental Go concepts like structs, functions, loops, and handling user input.

## Features

- Global Conferences: Book tickets for Go Conferences happening in different locations worldwide.
- Smart Validation: Ensures that the number of tickets you request is available and that your personal details (name, email) are valid.
- Live Updates: See the remaining tickets for each conference location in real-time.
- User-Friendly: The application guides you smoothly through the booking process with clear prompts.
- Instant Confirmation: Get immediate feedback after a successful booking, including details about your email confirmation.

## How to Run

To get this application up and running, you'll need Go installed on your system.

1. Clone the Repository:
      git clone https://github.com/TeamiruTesfahun/Go-Conference-Booking-CLI.git
   cd Go-Conference-Booking-CLI
   

2. Run the Application:
      go run main.go
   

## How to Use

1. Welcome: The application will greet you and show a list of countries where the Go Conference is being held.
2. Pick a Country: Enter the number corresponding to the country where you'd like to book tickets.
3. Conference Details: You'll see the location, date, and how many tickets are left for your chosen conference.
4. Enter Your Info: Provide your first name, last name, email address, and the number of tickets you want to buy.
5. Confirm Booking: If your input is correct and tickets are available, you'll see a confirmation message.
6. Book Again?: You'll then be asked if you want to book another ticket or exit the application.

## Code Structure

- main.go: This file contains all the core logic, including user interaction, input validation, and the booking process.
- Constants:
  - ConferenceName: Defines the name for the conference used in the application.
- Structs:
  - users: Holds details for each person who books a ticket, including their name, email, and number of tickets.
  - Country: Stores information about each conference location, such as its Name, Place, Date, and RemainingTickets.
- Data:
  - Countries: An array of Country structs, pre-loaded with all the available conference locations.
  - Booking: A slice of users structs that keeps track of all successful bookings.
- Functions:
  - GreetUsers(): Displays the welcoming message.
  - ValidateUserInput(): Checks if the entered names, email, and ticket quantity are valid.
  - BookTicket(): Manages the main booking process, updates the available tickets, and adds the user's booking to the Booking list.

## Future Enhancements

- Better Error Handling: Improve how the application handles invalid user input (e.g., if someone types text instead of a number for a choice).
- Data Storage: Implement a way to save booking information persistently, perhaps in a file or a database, so it's not lost when the application closes.
- Concurrency: Explore using Go's goroutines and channels to handle multiple bookings simultaneously.
- More User Data: Collect additional user information like phone numbers.
- Booking Cancellation: Add a feature allowing users to cancel their existing bookings.

## Contributing

Feel free to fork this repository, make your own improvements, and submit pull requests. All contributions are welcome!

## License

This project is open-source and available under the [MIT License](LICENSE).