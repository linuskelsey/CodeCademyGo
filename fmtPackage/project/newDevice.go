package main

import (
	"fmt"
)

func main() {
	var name string

	fmt.Print("What is your name? ")
	fmt.Scan(&name)

	fmt.Printf("Why hello there, %s!", name)

	var age int

	fmt.Print("\nHow old are you? ")
	fmt.Scan(&age)

	fmt.Printf("So let me get this straight - you're %s and you're %d years old!", name, age)

	newMessage := fmt.Sprintf("You're a %d-year old!, called %s", age, name)
	fmt.Println("\n" + newMessage)
}
