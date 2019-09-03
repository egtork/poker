package poker

import (
	"testing"
)

func TestNewCard(t *testing.T) {
	testCases := []struct {
		cardString string
		rank       Rank
		suit       Suit
	}{
		{"Ac", Ace, Clubs},
		{"Td", Ten, Diamonds},
		{"8h", Eight, Hearts},
		{"2s", Two, Spades},
	}
	for _, tc := range testCases {
		card := ParseCard(tc.cardString)
		if card.Rank != tc.rank {
			t.Errorf("Rank = %q, want %q", card.Rank, tc.rank)
		}
		if card.Suit != tc.suit {
			t.Errorf("Suit = %q, want %q", card.Suit, tc.suit)
		}
	}
}
