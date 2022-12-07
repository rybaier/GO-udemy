package main

import "fmt"

func main() {

	// cards := newDeck()

	// hand, remainingCards := deal(cards, 1)

	// hand.print()
	// remainingCards.print()

	cards := newDeck()
	fmt.Println(cards.toString())

}
