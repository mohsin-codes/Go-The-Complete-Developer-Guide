package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

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

func (d deck) toString() string {
	stringSlice := []string(d)
	stringToWrite := strings.Join(stringSlice, ", ")
	return stringToWrite
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func deckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return deck(strings.Split(string(bs), ", "))
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
