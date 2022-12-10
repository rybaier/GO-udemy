package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 { //Errorf can be changed to Error to avoid needing
		// formatting directive %v
		t.Errorf(" Expected deck length of 20, but got %v", len(d))

	}
}
