package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := newDeck()

	// if len(deck) < 52 {
	// 	t.Errorf("The expected number of cards are 52 and the cards are only %v", len(deck))
	// }

	if deck[0] != "Ace of Diamonds" {
		t.Errorf("The expected card is 'Ace of Diamonds', but the output is %v", deck[0])
	}

	if deck[len(deck)-1] != "Jack of Spades" {
		t.Errorf("The expected card is 'Jack of Spades', but the output is %v", deck[len(deck)-1])
	}

}

func TestSaveToFileAndDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	recoveredFile := deckFromFile("_decktesting")

	if len(recoveredFile) < 52 {
		t.Errorf("The expected length of deck is 52, but the length of the deck is %v", len(recoveredFile))
	}

	os.Remove("_decktesting")
}
