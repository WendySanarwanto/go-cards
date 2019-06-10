package main

import (
	"os"
	"testing"
)

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

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	filename := "_deckSaveLoadToFileTesting"
	os.Remove(filename)

	cards := newDeck()
	if cards.saveToFile(filename) != nil {
		t.Errorf("Expected deck to be saved into a file, named as %v", filename)
	}

	loadedCards, err := newDeckFromFile(filename)
	if err != nil {
		t.Errorf("Expected deck to be loaded successfully from '%v' file ", filename)
	}

	if (len(loadedCards) != len(cards)) {
		t.Errorf("Expected %v cards in loaded deck, but got %v cards", len(cards), len(loadedCards))
	}

	os.Remove(filename)
}