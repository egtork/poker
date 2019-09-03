package poker

import "fmt"

type Rank int

const (
	Two   Rank = 2
	Three Rank = 3
	Four  Rank = 4
	Five  Rank = 5
	Six   Rank = 6
	Seven Rank = 7
	Eight Rank = 8
	Nine  Rank = 9
	Ten   Rank = 10
	Jack  Rank = 11
	Queen Rank = 12
	King  Rank = 13
	Ace   Rank = 14
)

var rankToStringMap = map[Rank]string{
	Two:   "2",
	Three: "3",
	Four:  "4",
	Five:  "5",
	Six:   "6",
	Seven: "7",
	Eight: "8",
	Nine:  "9",
	Ten:   "T",
	Jack:  "J",
	Queen: "Q",
	King:  "K",
	Ace:   "A",
}

var parseRank = map[byte]Rank{
	'2': Two,
	'3': Three,
	'4': Four,
	'5': Five,
	'6': Six,
	'7': Seven,
	'8': Eight,
	'9': Nine,
	'T': Ten,
	'J': Jack,
	'Q': Queen,
	'K': King,
	'A': Ace,
}

func (r Rank) String() string {
	return rankToStringMap[r]
}

type Suit int

const (
	Clubs    Suit = iota
	Diamonds Suit = iota
	Hearts   Suit = iota
	Spades   Suit = iota
)

func (s Suit) String() string {
	if s == Clubs {
		return "c"
	} else if s == Diamonds {
		return "d"
	} else if s == Hearts {
		return "h"
	}
	return "s"
}

var parseSuit = map[byte]Suit{
	'c': Clubs,
	'd': Diamonds,
	'h': Hearts,
	's': Spades,
}

type Card struct {
	Rank Rank
	Suit Suit
}

func (c *Card) String() string {
	return fmt.Sprintf("%s%s", c.Rank, c.Suit)
}

func ParseCard(c string) *Card {
	r := parseRank[c[0]]
	s := parseSuit[c[1]]
	return &Card{r, s}
}

type Cards []*Card
