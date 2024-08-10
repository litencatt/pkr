package service

import (
	"github.com/litencatt/pkr/entity"
)

type PokerService interface {
	IsStartRound() bool
	StartRound() error
	GetRounds() int
	IsRoundWin() bool
	NextRound() error
	GetRoundStats() *entity.PokerRoundStats

	SelectCards([]string) error
	DrawCard(int) ([]entity.Trump, error)
	PlayHand() (entity.PokerHandStats, error)
	DiscardHand() error
	CancelHand() error

	GetCurrentAnteAmount() int
	GetCurrentBlindMulti() float64
	GetNextDrawNum() int
	GetChipAndMult(entity.HandType, int) (int, int)
	GetHandCardString() []string
	GetRemainCardString() []string
	GetEnableActions() []string
	SetAction(string)
}

type pokerService struct {
	config  PokerServiceConfig
	runInfo *entity.RunInfo
}

func NewPokerService(config PokerServiceConfig) PokerService {
	return &pokerService{
		config:  config,
		runInfo: entity.NewRunInfo(),
	}
}

type PokerServiceConfig struct {
	DebugMode bool
}

func (s *pokerService) GetNextDrawNum() int {
	if s.runInfo.Round.BeforeSelectAction == "" {
		return s.runInfo.DefaultDeal
	}

	if s.runInfo.Round.BeforeSelectAction == "Cancel" {
		return 0
	}

	return len(s.runInfo.Round.SelectedCards)
}

func (s *pokerService) GetChipAndMult(handType entity.HandType, level int) (int, int) {
	return s.runInfo.PokerHands.GetChipAndMult(handType, level)
}

func (s *pokerService) IsStartRound() bool {
	return s.runInfo.StartNext
}

func (s *pokerService) StartRound() error {
	s.runInfo.UnsetStartNext()

	scoreAtLeast := int(float64(s.GetCurrentAnteAmount()) * s.GetCurrentBlindMulti())
	s.runInfo.Round = entity.NewPokerRound(
		s.runInfo.Deck,
		s.runInfo.DefaultHands,
		s.runInfo.DefaultDiscards,
		scoreAtLeast,
	)
	s.runInfo.Round.Deck.Shuffle()

	return nil
}

func (s *pokerService) DrawCard(num int) ([]entity.Trump, error) {
	cards := s.runInfo.Round.DrawCard(num)
	return cards, nil
}

func (s *pokerService) GetCurrentAnteAmount() int {
	return entity.AnteAmounts[s.runInfo.AnteIndex]
}

func (s *pokerService) NextRound() error {
	s.runInfo.NextRound()
	scoreAtLeast := int(float64(s.GetCurrentAnteAmount()) * s.GetCurrentBlindMulti())
	s.runInfo.Round = entity.NewPokerRound(
		s.runInfo.Deck,
		s.runInfo.DefaultHands,
		s.runInfo.DefaultDiscards,
		scoreAtLeast,
	)
	return nil
}

func (s *pokerService) GetCurrentBlindMulti() float64 {
	return entity.BlindMultis[s.runInfo.BlindIndex]
}

func (s *pokerService) GetEnableActions() []string {
	var actions = []string{"Play"}
	if s.runInfo.Round.Discards > 0 {
		actions = append(actions, "Discard")
	}
	actions = append(actions, "Cancel")

	return actions
}

func (s *pokerService) SelectCards(cards []string) error {
	s.runInfo.Round.SetSelectCards(cards)

	return nil
}

func (s *pokerService) DiscardHand() error {
	s.runInfo.Discards--

	return nil
}

func (s *pokerService) CancelHand() error {
	s.runInfo.Round.SelectedCards = nil

	return nil
}

func (s *pokerService) PlayHand() (entity.PokerHandStats, error) {
	s.runInfo.Hands--

	round := s.runInfo.Round

	// get hand type and base chip and mult
	handType := round.PlayHand()
	chip, mult := s.GetChipAndMult(handType, 1)

	// get card rank and add to chip
	handsRankTotal := round.GetSelectCardsRankTotal()
	chip += handsRankTotal
	score := chip * mult
	round.TotalScore += score

	stats := entity.PokerHandStats{
		HandType: handType,
		Chip:     chip,
		Mult:     mult,
		Score:    score,
	}

	return stats, nil
}

func (s *pokerService) GetHandCardString() []string {
	return s.runInfo.Round.HandCardString()
}

func (s *pokerService) GetRemainCardString() []string {
	return s.runInfo.Round.RemainCardString()
}

func (s *pokerService) GetRoundStats() *entity.PokerRoundStats {
	return s.runInfo.Round.GetRoundStats()
}

func (s *pokerService) SetAction(action string) {
	s.runInfo.Round.BeforeSelectAction = action
}

func (s *pokerService) IsRoundWin() bool {
	return s.runInfo.Round.IsWin()
}

func (s *pokerService) GetRounds() int {
	return s.runInfo.Rounds
}

// NewPokerServiceConfig returns a new PokerServiceConfig
func NewPokerServiceConfig() PokerServiceConfig {
	return PokerServiceConfig{}
}
