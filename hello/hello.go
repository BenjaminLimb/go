package main

import (
	"fmt"
	"rsc.io/quote"
	"benjaminlimb.io/greetings"
)


func main (){
	fmt.Println("Hello, World!")

	fmt.Println(quote.Go())

	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
