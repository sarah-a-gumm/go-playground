package main

import (
	"fmt"

	"playground/greetings"
)

func main() {
	//Get a greeting message and print it.
	message := greetings.Hello("Sarah")
	fmt.Println(message)
}
