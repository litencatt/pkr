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
	DefaultDeal     int
	DefaultHands    int
	DefaultDiscards int
	AnteIndex       int
	BlindIndex      int
	Deck            Deck
	PokerHands      *PokerHands
	Rounds          int
	StartNext       bool
	JokerCards      []JokerCard
}

func NewRunInfo() *RunInfo {
	return &RunInfo{
		DefaultDeal:     8,
		DefaultHands:    4,
		DefaultDiscards: 3,
		Deck:            NewDeck(),
		PokerHands:      NewPokerHands(),
		Rounds:          1,
		StartNext:       true,
		AnteIndex:       0,
		BlindIndex:      0,
	}
}

func (r *RunInfo) UnsetStartNext() {
	r.StartNext = false
}

func (r *RunInfo) NextRound() error {
	r.Rounds += 1
	r.NextBlind()
	r.StartNext = true

	return nil
}

func (r *RunInfo) NextBlind() error {
	r.BlindIndex += 1
	if r.BlindIndex >= len(BlindMultis) {
		r.BlindIndex = 0
		r.NextAnte()
	}
	return nil
}

func (r *RunInfo) NextAnte() error {
	r.AnteIndex += 1
	return nil
}

func (s *RunInfo) AddJokerCard(jokerCard JokerCard) {
	s.JokerCards = append(s.JokerCards, jokerCard)
}

func (s *RunInfo) GetJokerCards() []JokerCard {
	return s.JokerCards
}
