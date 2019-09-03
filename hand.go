package poker

import "sort"

type Hand []*Card

func NewHandUnsorted(s string) Hand {
	if len(s)%2 != 0 {
		panic("NewHand input string must have even number of characters.")
	}
	n := len(s) / 2
	h := make(Hand, n, n)
	for i := 0; i < n; i++ {
		h[i] = ParseCard(s[2*i : 2*i+2])
	}
	return h
}

func NewHand(s string) Hand {
	h := NewHandUnsorted(s)
	h.Sort()
	return h
}

func (h Hand) Sort() {
	sort.SliceStable(h, func(i, j int) bool { return h[i].Rank > h[j].Rank })
}
