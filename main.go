package main

import "fmt"

func main() {
	cards, err := newDeckFromFile("MyCards")
	if err != nil {
		fmt.Println("Error: ", err)
		cards = newDeck()
		cards.saveToFile("MyCards")
	}
	// Shuffle the cards 1st
	cards.shuffle()
	cards.print()

	// Call deal method
	_, remainingDeck := deal(cards, 5)
	// hand.print()
	// remainingDeck.print()
	fmt.Println("Remaining Decks: ", remainingDeck.toString())
}
