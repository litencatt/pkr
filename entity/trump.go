package entity

import "sort"

type Suit string
type Rank string

const (
	Clubs    Suit = "Clubs"
	Diamonds      = "Diamonds"
	Hearts        = "Hearts"
	Spades        = "Spades"
)

const (
	Two   Rank = "2"
	Three      = "3"
	Four       = "4"
	Five       = "5"
	Six        = "6"
	Seven      = "7"
	Eight      = "8"
	Nine       = "9"
	Ten        = "T"
	Jack       = "J"
	Queen      = "Q"
	King       = "K"
	Ace        = "A"
)

type Trump struct {
	Suit Suit
	Rank Rank
}

func (t Trump) String() string {
	return string(t.Rank) + " of " + string(t.Suit)
}

func (t Trump) GetRankNumber() int {
	switch t.Rank {
	case Two:
		return 2
	case Three:
		return 3
	case Four:
		return 4
	case Five:
		return 5
	case Six:
		return 6
	case Seven:
		return 7
	case Eight:
		return 8
	case Nine:
		return 9
	case Ten, Jack, Queen, King:
		return 10
	case Ace:
		return 11
	}
	return 0
}

func (t Trump) GetSortOrder() int {
	switch t.Rank {
	case Two:
		return 2
	case Three:
		return 3
	case Four:
		return 4
	case Five:
		return 5
	case Six:
		return 6
	case Seven:
		return 7
	case Eight:
		return 8
	case Nine:
		return 9
	case Ten:
		return 10
	case Jack:
		return 11
	case Queen:
		return 12
	case King:
		return 13
	case Ace:
		return 14
	}
	return 0
}

func Contains(trumps []Trump, trump Trump) bool {
	for _, t := range trumps {
		if t == trump {
			return true
		}
	}
	return false
}

func Sort(trumps []Trump) {
	sort.Slice(trumps, func(i, j int) bool {
		return trumps[i].GetSortOrder() < trumps[j].GetSortOrder()
	})
}
