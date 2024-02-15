package main

import (
	"fmt"
	"math/rand"
	"rsc.io/quote"
)

func selectRandomQuote() string {
	quotesArray := []string{
		quote.Glass(),
		quote.Go(),
		quote.Hello(),
		quote.Opt(),
	}

	randomIndex := rand.Intn(len(quotesArray))
	return quotesArray[randomIndex]
}

func GreetUser() {
	for {
		fmt.Println("What's your name?")

		var inputName string
		randomQuote := selectRandomQuote()
		fmt.Scan(&inputName)

		if inputName == "exit" {
			fmt.Println("Exiting...")
			return
		}

		if inputName == "" {
			fmt.Println("Invalid input. Please enter a name.")
		} else {
			fmt.Printf("Hello, %s, here's your quote: %s", inputName, randomQuote)
			break
		}
	}
}

func main() {
	GreetUser()
}
