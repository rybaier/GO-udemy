package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 { //Errorf can be changed to Error to avoid needing
		// formatting directive %v
		t.Errorf(" Expected deck length of 52, but got %v", len(d))

	}
	if d[0] != "Ace of Spades" {
		t.Error("Expected first card to be Ace of Spades but got", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card to be King of Clubs but got %v", d[len(d)-1])
	}

}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")
	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards in deck but got %v", loadedDeck)
	}

	os.Remove("_decktesting")
}
