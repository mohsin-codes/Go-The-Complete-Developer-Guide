package main

import "fmt"

func main() {
	cards := newDeck()
	remainingDeck, hand := deal(cards, 5)
	hand.print()
	fmt.Println()
	remainingDeck.print()
}
