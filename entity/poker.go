package entity

var AnteAmounts = []int{
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

var BlindMultis = []float64{
	1.0,
	1.5,
	2.0,
}

type RunInfo struct {
	DefaultDeal int
	AnteIndex   int
	BlindIndex  int
	Deck        Deck
	PokerHands  *PokerHands
	Hands       int
	Discards    int
	Rounds      int
	Round       *PokerRound
}

func NewRunInfo() *RunInfo {
	return &RunInfo{
		DefaultDeal: 8,
		AnteIndex:   0,
		BlindIndex:  0,
		Deck:        NewDeck(),
		PokerHands:  NewPokerHands(),
		Hands:       4,
		Discards:    3,
		Rounds:      1,
		Round:       nil,
	}
}

func (r *RunInfo) NextAnte() error {
	r.AnteIndex += 1
	return nil
}

func (r *RunInfo) NextBlind() error {
	r.BlindIndex += 1
	if r.BlindIndex >= len(BlindMultis) {
		r.BlindIndex = 0

		r.NextRound()
		r.NextAnte()
	}
	return nil
}

func (r *RunInfo) NextRound() error {
	r.Rounds += 1
	return nil
}

func (r *RunInfo) ResetBlind() error {
	r.BlindIndex = 0
	return nil
}
