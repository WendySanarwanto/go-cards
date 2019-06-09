package main

import "fmt"

// Create a new Type of `deck`
// which is a slice of strings
type deck []string

// A helper method for creating an instance of deck type and initialise its elements
func newDeck() deck {
	cards := deck{}
	cardSuits := []string { "Clubs", "Diamonds", "Hearts", "Spades" }
	cardValues := 
		[]string { "Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", 
							 "Nine", "Ten", "Jack", "Queen", "King" }

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value + " of " + suit)
		}		
	}

	return cards
}

// Function with receiver
// It's like a method in class
func (_deck deck) print() {
	// Loop on each elements in the slice and print them
	for i, card := range _deck {
		fmt.Println("Card #", i+1, ": ", card)
	}
}
