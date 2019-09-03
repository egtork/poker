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
		desc: "StraightFlush",
		fn:   Hand.bestStraightFlushHand,
		p: []PositiveTestCase{
			PositiveTestCase{"2h3h4h5h6h", "6h5h4h3h2h"},
			PositiveTestCase{"Ah2h3h4h5h", "5h4h3h2hAh"},
			PositiveTestCase{"2h3h3d4h5h6h8s", "6h5h4h3h2h"},
		},
		n: []string{
			"2h3h4h5d6h7h8h",
			"2h3h4hTd6h7h8h",
		},
	},
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
