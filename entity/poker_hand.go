package entity

import (
	"sort"
)

type HandType string

const (
	HighCard      HandType = "High Card"
	OnePair       HandType = "One Pair"
	TwoPair       HandType = "Two Pair"
	ThreeOfAKind  HandType = "Three of a Kind"
	Straight      HandType = "Straight"
	Flush         HandType = "Flush"
	FullHouse     HandType = "Full House"
	FourOfAKind   HandType = "Four of a Kind"
	StraightFlush HandType = "Straight Flush"
	RoyalFlush    HandType = "Royal Flush"
)

type PokerHands struct {
	PokerHands []PokerHand
}

type PokerHand struct {
	HandType HandType
	Level    []PokerHandLevel
}

type PokerHandLevel struct {
	Level int
	Chip  int
	Mult  int
}

type PokerHandStats struct {
	HandType HandType
	Level    int
	Chip     int
	Mult     int
	Score    int
}

// FIXME: Set the correct chip and mult values for each hand type and level
func NewPokerHands() *PokerHands {
	return &PokerHands{
		PokerHands: []PokerHand{
			{
				HandType: HighCard,
				Level: []PokerHandLevel{
					{Level: 1, Chip: 5, Mult: 1},
					{Level: 2, Chip: 15, Mult: 2},
					{Level: 3, Chip: 15, Mult: 1},
					{Level: 4, Chip: 20, Mult: 1},
					{Level: 5, Chip: 25, Mult: 1},
					{Level: 6, Chip: 30, Mult: 1},
					{Level: 7, Chip: 35, Mult: 1},
					{Level: 8, Chip: 40, Mult: 1},
					{Level: 9, Chip: 45, Mult: 1},
					{Level: 10, Chip: 50, Mult: 1},
				},
			},
			{
				HandType: OnePair,
				Level: []PokerHandLevel{
					{Level: 1, Chip: 10, Mult: 2},
					{Level: 2, Chip: 4, Mult: 1},
					{Level: 3, Chip: 6, Mult: 1},
					{Level: 4, Chip: 8, Mult: 1},
					{Level: 5, Chip: 10, Mult: 1},
					{Level: 6, Chip: 12, Mult: 1},
					{Level: 7, Chip: 14, Mult: 1},
					{Level: 8, Chip: 16, Mult: 1},
					{Level: 9, Chip: 18, Mult: 1},
					{Level: 10, Chip: 20, Mult: 1},
				},
			},
			{
				HandType: TwoPair,
				Level: []PokerHandLevel{
					{Level: 1, Chip: 20, Mult: 2},
					{Level: 2, Chip: 6, Mult: 1},
					{Level: 3, Chip: 9, Mult: 1},
					{Level: 4, Chip: 12, Mult: 1},
					{Level: 5, Chip: 15, Mult: 1},
					{Level: 6, Chip: 18, Mult: 1},
					{Level: 7, Chip: 21, Mult: 1},
					{Level: 8, Chip: 24, Mult: 1},
					{Level: 9, Chip: 27, Mult: 1},
					{Level: 10, Chip: 30, Mult: 1},
				},
			},
			{
				HandType: ThreeOfAKind,
				Level: []PokerHandLevel{
					{Level: 1, Chip: 30, Mult: 3},
					{Level: 2, Chip: 2, Mult: 1},
					{Level: 3, Chip: 3, Mult: 1},
					{Level: 4, Chip: 4, Mult: 1},
					{Level: 5, Chip: 5, Mult: 1},
					{Level: 6, Chip: 6, Mult: 1},
					{Level: 7, Chip: 7, Mult: 1},
					{Level: 8, Chip: 8, Mult: 1},
					{Level: 9, Chip: 9, Mult: 1},
					{Level: 10, Chip: 10, Mult: 1}},
			},
			{
				HandType: Straight,
				Level: []PokerHandLevel{
					{Level: 1, Chip: 30, Mult: 4},
					{Level: 2, Chip: 2, Mult: 1},
					{Level: 3, Chip: 3, Mult: 1},
					{Level: 4, Chip: 4, Mult: 1},
					{Level: 5, Chip: 5, Mult: 1},
					{Level: 6, Chip: 6, Mult: 1},
					{Level: 7, Chip: 7, Mult: 1},
					{Level: 8, Chip: 8, Mult: 1},
					{Level: 9, Chip: 9, Mult: 1},
					{Level: 10, Chip: 10, Mult: 1}},
			},
			{
				HandType: Flush,
				Level: []PokerHandLevel{
					{Level: 1, Chip: 35, Mult: 4},
					{Level: 2, Chip: 2, Mult: 1},
					{Level: 3, Chip: 3, Mult: 1},
					{Level: 4, Chip: 4, Mult: 1},
					{Level: 5, Chip: 5, Mult: 1},
					{Level: 6, Chip: 6, Mult: 1},
					{Level: 7, Chip: 7, Mult: 1},
					{Level: 8, Chip: 8, Mult: 1},
					{Level: 9, Chip: 9, Mult: 1},
					{Level: 10, Chip: 10, Mult: 1}},
			},
			{
				HandType: FullHouse,
				Level: []PokerHandLevel{
					{Level: 1, Chip: 40, Mult: 4},
					{Level: 2, Chip: 2, Mult: 1},
					{Level: 3, Chip: 3, Mult: 1},
					{Level: 4, Chip: 4, Mult: 1},
					{Level: 5, Chip: 5, Mult: 1},
					{Level: 6, Chip: 6, Mult: 1},
					{Level: 7, Chip: 7, Mult: 1},
					{Level: 8, Chip: 8, Mult: 1},
					{Level: 9, Chip: 9, Mult: 1},
					{Level: 10, Chip: 10, Mult: 1}},
			},
			{
				HandType: FourOfAKind,
				Level: []PokerHandLevel{
					{Level: 1, Chip: 60, Mult: 7},
					{Level: 2, Chip: 2, Mult: 1},
					{Level: 3, Chip: 3, Mult: 1},
					{Level: 4, Chip: 4, Mult: 1},
					{Level: 5, Chip: 5, Mult: 1},
					{Level: 6, Chip: 6, Mult: 1},
					{Level: 7, Chip: 7, Mult: 1},
					{Level: 8, Chip: 8, Mult: 1},
					{Level: 9, Chip: 9, Mult: 1},
					{Level: 10, Chip: 10, Mult: 1}},
			},
			{
				HandType: StraightFlush,
				Level: []PokerHandLevel{
					{Level: 1, Chip: 100, Mult: 8},
					{Level: 2, Chip: 2, Mult: 1},
					{Level: 3, Chip: 3, Mult: 1},
					{Level: 4, Chip: 4, Mult: 1},
					{Level: 5, Chip: 5, Mult: 1},
					{Level: 6, Chip: 6, Mult: 1},
					{Level: 7, Chip: 7, Mult: 1},
					{Level: 8, Chip: 8, Mult: 1},
					{Level: 9, Chip: 9, Mult: 1},
					{Level: 10, Chip: 10, Mult: 1}},
			},
			{
				HandType: RoyalFlush,
				Level: []PokerHandLevel{
					{Level: 1, Chip: 100, Mult: 8},
					{Level: 2, Chip: 2, Mult: 1},
					{Level: 3, Chip: 3, Mult: 1},
					{Level: 4, Chip: 4, Mult: 1},
					{Level: 5, Chip: 5, Mult: 1},
					{Level: 6, Chip: 6, Mult: 1},
					{Level: 7, Chip: 7, Mult: 1},
					{Level: 8, Chip: 8, Mult: 1},
					{Level: 9, Chip: 9, Mult: 1},
					{Level: 10, Chip: 10, Mult: 1}},
			},
		},
	}
}

func (p *PokerHands) GetChipAndMult(HandType HandType, Level int) (Chip int, Mult int) {
	for _, ph := range p.PokerHands {
		if ph.HandType == HandType {
			for _, lvl := range ph.Level {
				if lvl.Level == Level {
					return lvl.Chip, lvl.Mult
				}
			}
		}
	}
	return 0, 0
}

// isFlush checks if all cards in the hand have the same suit.
func isFlush(hand []Trump) bool {
	if len(hand) < 5 {
		return false
	}

	firstSuit := hand[0].Suit
	for _, card := range hand[1:] {
		if card.Suit != firstSuit {
			return false
		}
	}
	return true
}

// isStraight checks if the hand forms a sequential ranking.
func isStraight(hand []Trump) bool {
	if len(hand) < 5 {
		return false
	}

	// Sort the hand by rank (implement the sorting logic based on your Rank definition)
	sort.Slice(hand, func(i, j int) bool {
		// Implement the comparison logic for your Rank type
		return hand[i].Rank < hand[j].Rank
	})

	// Check for Ace-low straight
	if hand[0].Rank == Two && hand[1].Rank == Three && hand[2].Rank == Four && hand[3].Rank == Five && hand[4].Rank == Ace {
		return true
	}

	// Check for standard straight
	for i := 0; i < len(hand)-1; i++ {
		// convert Rank string to int
		rank1 := hand[i].GetRankNumber()
		rank2 := hand[i+1].GetRankNumber()
		if rank2-rank1 != 1 {
			return false
		}
	}
	return true
}

// isRoyalFlush checks if the hand is a Royal Flush (10, J, Q, K, A).
func isRoyalFlush(hand []Trump) bool {
	if len(hand) != 5 {
		return false
	}
	
	requiredRanks := map[Rank]bool{
		Ten:   false,
		Jack:  false,
		Queen: false,
		King:  false,
		Ace:   false,
	}
	
	for _, card := range hand {
		if _, exists := requiredRanks[card.Rank]; exists {
			requiredRanks[card.Rank] = true
		} else {
			return false
		}
	}
	
	// Check if all required ranks are present
	for _, present := range requiredRanks {
		if !present {
			return false
		}
	}
	
	return true
}

// groupByRank groups cards by their ranks and returns a map of rank to count.
func groupByRank(hand []Trump) map[Rank]int {
	rankCount := make(map[Rank]int)
	for _, card := range hand {
		rankCount[card.Rank]++
	}
	return rankCount
}

// EvaluateHand evaluates the given hand and returns the HandType.
func EvaluateHand(hand []Trump) HandType {
	isFlush := isFlush(hand)
	isStraight := isStraight(hand)

	if isFlush && isStraight {
		// Check for Royal Flush (10, J, Q, K, A)
		if isRoyalFlush(hand) {
			return RoyalFlush
		}
		return StraightFlush
	}

	rankCount := groupByRank(hand)
	var pairs, threes, fours int
	for _, count := range rankCount {
		switch count {
		case 2:
			pairs++
		case 3:
			threes++
		case 4:
			fours++
		}
	}

	if fours == 1 {
		return FourOfAKind
	} else if threes == 1 && pairs == 1 {
		return FullHouse
	} else if isFlush {
		return Flush
	} else if isStraight {
		return Straight
	} else if threes == 1 {
		return ThreeOfAKind
	} else if pairs == 2 {
		return TwoPair
	} else if pairs == 1 {
		return OnePair
	}

	return HighCard
}

func GetScore(hand HandType) int {
	switch hand {
	case HighCard:
		return 1
	case OnePair:
		return 2
	case TwoPair:
		return 3
	case ThreeOfAKind:
		return 4
	case Straight:
		return 5
	case Flush:
		return 6
	case FullHouse:
		return 7
	case FourOfAKind:
		return 8
	case StraightFlush:
		return 9
	case RoyalFlush:
		return 10
	}
	return 0
}
