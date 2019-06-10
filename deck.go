package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Create a new Type of `deck`
// which is a slice of strings
type deck []string

// Read saved deck from a file
func newDeckFromFile(filename string) (deck, error) {
	readByteSlice, err := ioutil.ReadFile(filename)
	readData := string(readByteSlice)
	return deck(strings.Split(readData, ", ")), err
}

// Save joined deck's cards into a file
func (_deck deck) saveToFile(filename string) error {
	deckInBytes := []byte(_deck.toString())
	err := ioutil.WriteFile(filename, deckInBytes, 0666)
	return err
}

// Joined all cards in deck as a single string
func (_deck deck) toString() string {
	cards := []string(_deck);
	joinedCards := strings.Join(cards, ", ")
	return joinedCards
}

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

// Deak cards - Demonstrates a function which return multiple values
// 							and getting partial elements of slices by using range index
func deal(_deck deck, handSize int) (deck, deck) {
	return _deck[handSize:], // Pick element at specified index up to beyond
				 _deck[:handSize]  // Pick elements from index 0 up to specified index - 1
}
