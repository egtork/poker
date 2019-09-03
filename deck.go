package poker

import (
	"errors"
	"math/rand"
)

const (
	CardsPerDeck = 52
)

var suits = []Suit{Hearts, Spades, Clubs, Diamonds}
var ranks = []Rank{Ace, King, Queen, Jack, Ten, Nine, Eight, Seven, Six, Five, Four, Three, Two}

type Deck struct {
	i     int
	Cards Cards
}

func NewDeck() *Deck {
	cards := make(Cards, 0, CardsPerDeck)
	for _, s := range suits {
		for _, r := range ranks {
			cards = append(cards, &Card{Rank: r, Suit: s})
		}
	}
	d := &Deck{0, cards}
	d.Shuffle()
	return d
}

func (d *Deck) Shuffle() {
	d.i = 0
	rand.Shuffle(len(d.Cards), func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
}

func (d *Deck) Deal() (error, *Card) {
	if d.i == CardsPerDeck {
		return errors.New("No more cards in deck."), nil
	}
	c := d.Cards[d.i]
	d.i++
	return nil, c
}
