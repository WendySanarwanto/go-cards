package main

func main() {
	cards := newDeck()
	// cards.print()

	// Call deal method
	hand, remainingDeck := deal(cards, 5)
	hand.print()
	remainingDeck.print()
}
