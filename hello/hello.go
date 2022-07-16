//Declare a main package. In Go, code executed as an application must be in a main package.
package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	//A slice of names
	names := []string{"Gladys", "Samantha", "Darrin"}

	// Get a greeting message and print it.
	message, err := greetings.Hellos(names)

	if err != nil {
		// If an error was returned, print it to the console and
		// exit the program.
		log.Fatal(err)
	}

	fmt.Println(message)
}
