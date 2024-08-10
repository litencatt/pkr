package entity

var Antes = []int{
	300,
	800,
	2800,
	6000,
	11000,
	20000,
	35000,
	50000,
	110000,
	560000,
	7200000,
	300000000,
	47000000000,
	2900 * 100000000000,
	7700 * 1000000000000,
	// FIXME: return 8600e20
	8600,
}

// Multi
var Blinds = []float64{
	1.0,
	1.5,
	2.0,
}

type RunInfo struct {
	DefaultDeal int
	Ante        int
	Blind       float64
	Deck        Deck
	PokerHands  *PokerHands
	Hands       int
	Discards    int
	Round       *PokerRound
}

func NewRunInfo() *RunInfo {
	return &RunInfo{
		DefaultDeal: 8,
		Ante:        Antes[0],
		Blind:       Blinds[0],
		Deck:        NewDeck(),
		PokerHands:  NewPokerHands(),
		Hands:       4,
		Discards:    3,
		Round:       nil,
	}
}
