package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings

type deck []string

// d deck is the reciever
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Hearts", "Spades", "Clubs", "Diamonds"}

	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}

	return cards
}
