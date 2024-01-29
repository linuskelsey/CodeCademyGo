/*
	Unix Nano (Go) docs: https://pkg.go.dev/time#Time.UnixNano

	In random.go, the random number supplied for amountLeft is always the same. This is since the number is chosen using a 'seed', or a start
	point for its randomisations. By default that value is 1, and so we need to choose a new seed dby 'reseeding' before using rand.Intn().
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Add your code below:
	rand.Seed(time.Now().UnixNano())

	amountLeft := rand.Intn(10000)

	fmt.Println("amountLeft is: ", amountLeft)

	if amountLeft > 5000 {
		fmt.Println("What should I spend this on?")
	} else {
		fmt.Println("Where did all my money go?")
	}
}
