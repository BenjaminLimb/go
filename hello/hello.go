package main

import (
	"fmt"
	"log"
	"rsc.io/quote"
	"benjaminlimb.io/greetings"
)
func valid(){
	message, err := greetings.Hello("Gladys")
	// If an error was returned, print it to the console and exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message to the console.
	fmt.Println(message)
}

func invalid(){
	message, err := greetings.Hello("")
	// If an error was returned, print it to the console and exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message to the console.
	fmt.Println(message)
}
func main (){
	// Set properties of the predefined Logger,
	// including the log entry prefix and a flag to disable print the time,
	// source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	fmt.Println("Hello, World!")

	fmt.Println(quote.Go())

	log.Println("Call greetings.Hello with \"Gladys\"")
	valid()

	log.Println("Call greetings.Hello with \"\"")
	invalid()

}
