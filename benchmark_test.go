package poker

import (
	"testing"
)

const (
	cardsPerHand   = 5
	handsGenerated = 100000
)

func setUpHands(m int, sort bool) []Hand {
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
			if sort {
				hand.Sort()
			}
			hands = append(hands, hand)
			i++
		}
	}
	return hands
}

var result Hand
var sortedHands = setUpHands(handsGenerated, true)

func BenchmarkBestHand(b *testing.B) {
	sort := false
	hands := setUpHands(handsGenerated, sort)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, result = hands[n%handsGenerated].BestHand()
	}
}

func BenchmarkBestPairHand(b *testing.B) {
	for n := 0; n < b.N; n++ {
		result = sortedHands[n%handsGenerated].bestPairHand()
	}
}

func BenchmarkBestTwoPairHand(b *testing.B) {
	for n := 0; n < b.N; n++ {
		result = sortedHands[n%handsGenerated].bestTwoPairHand()
	}
}

func BenchmarkBestThreeOfAKindHand(b *testing.B) {
	for n := 0; n < b.N; n++ {
		result = sortedHands[n%handsGenerated].bestThreeOfAKindHand()
	}
}

func BenchmarkBestStraightHand(b *testing.B) {
	for n := 0; n < b.N; n++ {
		result = sortedHands[n%handsGenerated].bestStraightHand()
	}
}

func BenchmarkBestFlushHand(b *testing.B) {
	for n := 0; n < b.N; n++ {
		result = sortedHands[n%handsGenerated].bestThreeOfAKindHand()
	}
}

func BenchmarkBestFullHouseHand(b *testing.B) {
	for n := 0; n < b.N; n++ {
		result = sortedHands[n%handsGenerated].bestFullHouseHand()
	}
}

func BenchmarkBestFourOfAKindHand(b *testing.B) {
	for n := 0; n < b.N; n++ {
		result = sortedHands[n%handsGenerated].bestFourOfAKindHand()
	}
}

func BenchmarkBestStraightFlushHand(b *testing.B) {
	for n := 0; n < b.N; n++ {
		result, _ = sortedHands[n%handsGenerated].bestStraightFlushHand()
	}
}
