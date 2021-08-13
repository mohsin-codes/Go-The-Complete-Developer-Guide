package main

import "fmt"

//Create a new type of "deck"
//which is a slice of strings
type deck []string

//newDeck Function
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Diamonds", "Clubs", "Hearts", "Spades"}
	cardNumbers := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "King", "Queen", "Jack"}

	for _, suit := range cardSuits {
		for _, number := range cardNumbers {
			card := number + " of " + suit
			cards = append(cards, card)
		}
	}
	return cards
}

//reciever function
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	hand := d[:handSize]
	d = d[handSize:]
	return d, hand
}
