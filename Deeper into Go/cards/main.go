package main

import "fmt"

func main() {
	cards := newDeck()
	fmt.Println("Using Print Function:")
	cards.print()
	deal(cards, 5)
	fmt.Println()
	fmt.Println("Using Print after Deal Function:")
	cards.print()
	cards.saveToFile("newDeck.txt")
	fmt.Println()
	fmt.Println("Using Print Function After save to file and deckfromFile Function:")
	deckFromFile("newDeck.txt").print()
	cards.shuffle()
	fmt.Println()
	fmt.Println("Using Print Function After Shuffle Function:")
	cards.print()
}
