package poker

import (
	"testing"
)

const (
	cardsPerHand   = 8
	handsGenerated = 100000
)

var result Hand

func setUpHands(m int) []Hand {
	hands := make([]Hand, 0, m)
	i := 0
	for i < handsGenerated {
		deck := NewDeck()
		for j := 0; j < 52/cardsPerHand; j++ {
			hand := make(Hand, 0, cardsPerHand)
			for k := 0; k < cardsPerHand; k++ {
				_, card := deck.Deal()
				hand = append(hand, card)
			}
			hands = append(hands, hand)
			i++
		}
	}
	return hands
}

func BenchmarkBestHand(b *testing.B) {
	hands := setUpHands(handsGenerated)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, result = hands[n%handsGenerated].BestHand()
	}
}
