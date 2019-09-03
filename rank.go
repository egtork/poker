package poker

import "log"

type HandRank int

const (
	HighCard      HandRank = iota
	Pair          HandRank = iota
	TwoPair       HandRank = iota
	ThreeOfAKind  HandRank = iota
	Straight      HandRank = iota
	Flush         HandRank = iota
	FullHouse     HandRank = iota
	FourOfAKind   HandRank = iota
	StraightFlush HandRank = iota
)

func (h Hand) bestStraightFlushHand() Hand {
	suited := h.getFlushHand(5, 7)
	if suited != nil {
		return suited.bestStraightHand()
	}
	return nil
}

func (h Hand) bestFourOfAKindHand() Hand {
	for i := 0; i < len(h)-3; i++ {
		if h[i].Rank == h[i+1].Rank &&
			h[i+1].Rank == h[i+2].Rank &&
			h[i+2].Rank == h[i+3].Rank {
			var kicker *Card
			if i == 0 {
				kicker = h[4]
			} else {
				kicker = h[0]
			}
			return append(h[i:i+4], kicker)
		}
	}
	return nil
}

func (h Hand) bestFullHouseHand() Hand {
	var trips Cards
	for i := 0; i < len(h)-2; i++ {
		if h[i].Rank == h[i+1].Rank && h[i+1].Rank == h[i+2].Rank {
			trips = Cards{h[i], h[i+1], h[i+2]}
			break
		}
	}
	if trips == nil {
		return nil
	}
	for i := 0; i < len(h)-1; i++ {
		if h[i].Rank == h[i+1].Rank &&
			h[i] != trips[0] && h[i] != trips[1] && h[i] != trips[2] {
			return Hand(append(trips, Cards{h[i], h[i+1]}...))
		}
	}
	return nil
}

func (h Hand) bestFlushHand() Hand {
	return h.getFlushHand(5, 5)
}

func (h Hand) getFlushHand(min, max int) Hand {
	clubs, diamonds, hearts, spades := 0, 0, 0, 0
	for _, card := range h {
		switch card.Suit {
		case Clubs:
			clubs++
		case Diamonds:
			diamonds++
		case Hearts:
			hearts++
		case Spades:
			spades++
		}
	}
	var flushSuit Suit
	var flushFound bool
	if clubs >= min {
		flushSuit = Clubs
		flushFound = true
	} else if diamonds >= min {
		flushSuit = Diamonds
		flushFound = true
	} else if hearts >= min {
		flushSuit = Hearts
		flushFound = true
	} else if spades >= min {
		flushSuit = Spades
		flushFound = true
	}
	if flushFound {
		result := make(Hand, 0, max)
		for _, card := range h {
			if card.Suit == flushSuit {
				result = append(result, card)
				if len(result) == max {
					break
				}
			}
		}
		return result
	}
	return nil
}

func (h Hand) bestStraightHand() Hand {
	var run Hand
	i := 0
	for i < len(h)-3 {
		run = h[i : i+1]
		for i < len(h)-1 {
			i++
			if h[i-1].Rank == h[i].Rank+1 {
				run = append(run, h[i])
				if len(run) == 5 {
					return run
				}
			} else if h[i-1].Rank == h[i].Rank {
			} else {
				break
			}
		}
	}
	if h[0].Rank == Ace &&
		len(run) == 4 && run[0].Rank == Five && run[3].Rank == Two {
		return append(run, h[0])
	}
	return nil
}

func (h Hand) bestThreeOfAKindHand() Hand {
	for i := 0; i < len(h)-2; i++ {
		if h[i].Rank == h[i+1].Rank && h[i+1].Rank == h[i+2].Rank {
			if i == 0 {
				return h[:5]
			} else if i == 1 {
				return Hand{h[1], h[2], h[3], h[0], h[4]}
			}
			return append(h[i:i+3], h[0:2]...)
		}
	}
	return nil
}

func (h Hand) bestTwoPairHand() Hand {
	var firstPair Cards
	var secondPair Cards
	var kicker *Card
	for i := 0; i < len(h); i++ {
		if i < len(h)-1 && h[i] == h[i+1] {
			if firstPair == nil {
				firstPair = Cards{h[i], h[i+1]}
			} else {
				secondPair = Cards{h[i], h[i+1]}
			}
			i += 1
		} else if kicker == nil {
			kicker = h[i]
		}
		if firstPair != nil && secondPair != nil && kicker != nil {
			pairs := append(firstPair, secondPair...)
			return Hand(append(pairs, kicker))
		}
	}
	return nil
}

func (h Hand) bestPairHand() Hand {
	for i := 0; i < len(h)-1; i++ {
		if h[i].Rank == h[i+1].Rank {
			if i == 0 {
				return h[0:5]
			} else if i == 1 {
				return Hand{h[1], h[2], h[0], h[3], h[4]}
			} else if i == 2 {
				return Hand{h[2], h[3], h[0], h[1], h[4]}
			}
			return append(h[i:i+2], h[0:3]...)
		}
	}
	return nil
}

func (h Hand) bestHighCardHand() Hand {
	return h[:5]
}

func (h Hand) BestHand() (HandRank, Hand) {
	if len(h) < 5 || len(h) > 7 {
		log.Panicf("Invalid hand size: %d\n", len(h))
	}
	h.Sort()
	if hand := h.bestStraightFlushHand(); hand != nil {
		return StraightFlush, hand
	} else if hand := h.bestFourOfAKindHand(); hand != nil {
		return FourOfAKind, hand
	} else if hand := h.bestFullHouseHand(); hand != nil {
		return FullHouse, hand
	} else if hand := h.bestFlushHand(); hand != nil {
		return Flush, hand
	} else if hand := h.bestStraightHand(); hand != nil {
		return Straight, hand
	} else if hand := h.bestThreeOfAKindHand(); hand != nil {
		return ThreeOfAKind, hand
	} else if hand := h.bestTwoPairHand(); hand != nil {
		return TwoPair, hand
	} else if hand := h.bestPairHand(); hand != nil {
		return Pair, hand
	}
	return HighCard, h.bestHighCardHand()
}
