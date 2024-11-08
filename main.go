package main

import (
	"fmt"
	"os"
)

func main() {
	var cardNumber int
	fmt.Printf("Enter card number: ")
	_, err := fmt.Scanln(&cardNumber)
	if err != nil {
		fmt.Println("Something went wrong")
		os.Exit(1)
	}

	if isCardValid(cardNumber) {
		fmt.Println("Correct")
	} else {
		fmt.Println("Incorrect")
	}
}
