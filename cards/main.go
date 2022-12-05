package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	// define variable in multiple ways with identifying type or
	// using :=
	// can redefine variable with =
	// can initialize variable without assigning a value
	// can initialize variable outside of function without assigning value
	// can only assign value to variable within function
	cards := []string{newCard(), "Ace of Diamonds"}
	cards = append(cards, "Six of Spades")

	for i, card := range cards {
		fmt.Println(i, card)
	}
	// fmt.Println(cards)

}

func newCard() string {
	return "Five of Diamonds"
}
