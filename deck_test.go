package main

import "testing"

func TestNewDeck(t *testing.T) {
	cards := newDeck()
	cardsLength := len(cards)
	expectedCardsLength := 52
	if (cardsLength != expectedCardsLength) {
		t.Errorf("Expected number of cards is %v, but got %v", 
							expectedCardsLength, cardsLength)
	}

	if (cards[0] != "Ace of Clubs") {
		t.Errorf("Expected card #1 is '%v', but got '%v'", "Ace of Clubs", cards[0])
	}

	if (cards[expectedCardsLength-1] != "King of Spades") {
		t.Errorf("Expected card #%v is '%v', but got '%v'", 
							cards[expectedCardsLength-1],  "King of Spades", cards[0])
	}
}