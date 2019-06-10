package main

import "fmt"

func main() {
	cards := newDeck()
	// cards.print()

	// Call deal method
	_, remainingDeck := deal(cards, 5)
	// hand.print()
	// remainingDeck.print()
	fmt.Println("Remaining Decks: ", remainingDeck.toString())
}
