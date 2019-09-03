package poker

import (
	"testing"
)

func cardsEqual(h1, h2 Hand) bool {
	if h1 == nil && h2 == nil {
		return true
	} else if h1 == nil || h2 == nil {
		return false
	} else if len(h1) != len(h2) {
		return false
	}
	for i, c := range h1 {
		if c.Rank != h2[i].Rank {
			return false
		}
		if c.Suit != h2[i].Suit {
			return false
		}
	}
	return true
}

func TestNewHand(t *testing.T) {
	h := NewHand("AsJsJd8c6d3s2h")
	wantLen := 7
	if len(h) != wantLen {
		t.Errorf("Len = %q, want %q", len(h), wantLen)
	}
}

type HandRankingFn func(Hand) Hand

type PositiveTestCase struct {
	in   string
	want string
}

type Test struct {
	desc string
	fn   HandRankingFn
	p    []PositiveTestCase
	n    []string
}

var tests = []Test{
	Test{
		desc: "FourOfAKind",
		fn:   Hand.bestFourOfAKindHand,
		p: []PositiveTestCase{
			PositiveTestCase{"AhAsKhAcAd", "AhAsAcAdKh"},
			PositiveTestCase{"Ah9hAsKhAc8dAd", "AhAsAcAdKh"},
		},
		n: []string{
			"5h4h5d4d5s",
			"5h5d5s4h4d4s3h",
		},
	},
	Test{
		desc: "FullHouse",
		fn:   Hand.bestFullHouseHand,
		p: []PositiveTestCase{
			PositiveTestCase{"4h4d4s5h5d5s5c", "5h5d5s4h4d"},
			PositiveTestCase{"AhAdAsKhKd", "AhAdAsKhKd"},
			PositiveTestCase{"AhAdKsKhKd", "KsKhKdAhAd"},
		},
		n: []string{
			"AhAdAsKhQh",
			"AhAdKhKdQhQdJh",
		},
	},
	Test{
		desc: "Flush",
		fn:   Hand.bestFlushHand,
		p: []PositiveTestCase{
			PositiveTestCase{"4h6h8hThQh", "QhTh8h6h4h"},
			PositiveTestCase{"AdQhKdTd3d2h5d", "AdKdTd5d3d"},
			PositiveTestCase{"9s8s7sKsQsJsTs", "KsQsJsTs9s"},
		},
		n: []string{
			"4h5hTs6h8h",
			"2d4d6d9s8s7sKs",
		},
	},
	Test{
		desc: "Straight",
		fn:   Hand.bestStraightHand,
		p: []PositiveTestCase{
			PositiveTestCase{"2h3h4h5d6d", "6d5d4h3h2h"},
			PositiveTestCase{"2h3h4h5dAd", "5d4h3h2hAd"},
			PositiveTestCase{"AdKdQdJdTh", "AdKdQdJdTh"},
			PositiveTestCase{"AdQdJdTh9s8s3s", "QdJdTh9s8s"},
		},
		n: []string{
			"2h3h4h5d7d",
			"AdQdJd7h9s8s3s",
		},
	},
	Test{
		desc: "ThreeOfAKind",
		fn:   Hand.bestThreeOfAKindHand,
		p: []PositiveTestCase{
			PositiveTestCase{"AsAh8cJh3s2hAd", "AsAhAdJh8c"},
			PositiveTestCase{"AsJsJd8cJh3s2h", "JsJdJhAs8c"},
			PositiveTestCase{"2sJs2d8c9h3s2h", "2s2d2hJs9h"},
		},
		n: []string{
			"8cAsJs6d3h",
		},
	},
	Test{
		desc: "TwoPair",
		fn:   Hand.bestTwoPairHand,
		p: []PositiveTestCase{
			PositiveTestCase{"AsJsJd8c6d3s6s", "JsJd6d6sAs"},
			PositiveTestCase{"AsJsJdJc6dAc2h", "AsAcJsJdJc"},
			PositiveTestCase{"6h2h3h2c6c", "6h6c2h2c3h"},
		},
		n: []string{
			"8cAsJs6d3h",
		},
	},
	Test{
		desc: "Pair",
		fn:   Hand.bestPairHand,
		p: []PositiveTestCase{
			PositiveTestCase{"AsJsJd8c6d3s2h", "JsJdAs8c6d"},
			PositiveTestCase{"AsJsJdJc6d3s2h", "JsJdAsJc6d"},
			PositiveTestCase{"AsJsTd9c6d2s2h", "2s2hAsJsTd"},
			PositiveTestCase{"AsJs8c6d3sAh", "AsAhJs8c6d"},
		},
		n: []string{
			"8cAsJs6d3h",
		},
	},
	Test{
		desc: "HighCard",
		fn:   Hand.bestHighCardHand,
		p: []PositiveTestCase{
			PositiveTestCase{"KhQdTc9d7s", "KhQdTc9d7s"},
			PositiveTestCase{"KhQdTc9d7s5d4s", "KhQdTc9d7s"},
		},
		n: []string{},
	},
}

func TestRanking(t *testing.T) {
	for _, test := range tests {
		for _, pTest := range test.p {
			got := test.fn(NewHand(pTest.in))
			want := NewHandUnsorted(pTest.want)
			if !cardsEqual(got, want) {
				t.Errorf("%v: Got %v, want %v", test.desc, got, want)
			}
		}
		for _, nTest := range test.n {
			if got := test.fn(NewHand(nTest)); got != nil {
				t.Errorf("%v: Got %v, want nil.", test.desc, got)
			}
		}
	}
}

func TestStraighFlushRanking(t *testing.T) {
	// Straight flush
	t1 := [][]string{
		[]string{"2h3h4h5h6h", "6h5h4h3h2h", "6h5h4h3h2h"},
		[]string{"Ah2h3h4h5h", "5h4h3h2hAh", "Ah5h4h3h2h"},
		[]string{"2h3h3d4h5h6hAh", "6h5h4h3h2h", "Ah6h5h4h3h"},
	}
	for _, test := range t1 {
		sfGot, fGot := Hand.bestStraightFlushHand(NewHand(test[0]))
		sfWant := NewHandUnsorted(test[1])
		fWant := NewHandUnsorted(test[2])
		if !cardsEqual(sfGot, sfWant) {
			t.Errorf("Straight flush: Got %v, want %v", sfGot, sfWant)
		}
		if !cardsEqual(fGot, fWant) {
			t.Errorf("Flush: Got %v, want %v", fGot, fWant)
		}
	}
	// No straight flush, flush
	t2 := [][]string{
		[]string{"2h3h4h5h7h", "7h5h4h3h2h"},
		[]string{"AhKh3h4h5h", "AhKh5h4h3h"},
		[]string{"AsKhJs9h8h7h6h", "Kh9h8h7h6h"},
	}
	for _, test := range t2 {
		sfGot, fGot := Hand.bestStraightFlushHand(NewHand(test[0]))
		fWant := NewHandUnsorted(test[1])
		if sfGot != nil {
			t.Errorf("Got %v, want nil.", sfGot)
		}
		if !cardsEqual(fGot, fWant) {
			t.Errorf("Got %v, want %v", fGot, fWant)
		}
	}
	// No straight flush, no flush
	t3 := []string{
		"2h3d4h5d6h7h8s",
		"2h3d4hTd6h7h8s",
	}
	for _, test := range t3 {
		sfGot, fGot := Hand.bestStraightFlushHand(NewHand(test))
		if sfGot != nil {
			t.Errorf("Got %v, want nil.", sfGot)
		}
		if fGot != nil {
			t.Errorf("Got %v, want nil", fGot)
		}
	}
}
