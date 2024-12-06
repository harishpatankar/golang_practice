package main

import (
	"fmt"

	greetings "github.com/harishpatankar/golang_practice/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
