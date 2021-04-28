package main

import (
	"fmt"
	"log"
	"rsc.io/quote"
	"benjaminlimb.io/greetings"
)
func greet(name string){
	message, err := greetings.Hello(name)
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
	greet("Gladys")

	// A slice of names.
	names := []string{"Benjamin", "Madison", "Joseph", "Matthew"}
	messages, err := greetings.Hellos(names)
	if err !=nil{
		log.Fatal(err)
	}

	fmt.Println(messages)

	log.Println("Call greetings.Hello with \"\"")
	greet("")

}
