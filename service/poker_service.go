package service

import (
	"fmt"

	"github.com/litencatt/pkr/entity"
)

type PokerService interface {
	IsStartRound() bool
	StartRound([]entity.JokerCard) error
	GetRounds() int
	IsRoundWin() bool
	NextRound() error
	GetRoundStats() *entity.RoundStats

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

	ShopOpen() []string
	AddShopItem(string) error

	ShowJokers()
}

type pokerService struct {
	config  PokerServiceConfig
	runInfo *entity.RunInfo
	round   *entity.PokerRound
}

func NewPokerService(config PokerServiceConfig) PokerService {
	runInfo := entity.NewRunInfo()
	round := entity.NewPokerRound(
		runInfo.Deck,
		runInfo.DefaultHands,
		runInfo.DefaultDiscards,
		runInfo.DefaultDeal,
	)

	return &pokerService{
		config:  config,
		runInfo: runInfo,
		round:   round,
	}
}

type PokerServiceConfig struct {
	DebugMode bool
}

func (s *pokerService) GetNextDrawNum() int {
	if s.round.BeforeSelectAction == "" {
		return s.runInfo.DefaultDeal
	}

	if s.round.BeforeSelectAction == "Cancel" {
		return 0
	}

	return len(s.round.SelectedCards)
}

func (s *pokerService) GetChipAndMult(handType entity.HandType, level int) (int, int) {
	return s.runInfo.PokerHands.GetChipAndMult(handType, level)
}

func (s *pokerService) IsStartRound() bool {
	return s.runInfo.StartNext
}

func (s *pokerService) StartRound([]entity.JokerCard) error {
	s.runInfo.UnsetStartNext()

	scoreAtLeast := int(float64(s.GetCurrentAnteAmount()) * s.GetCurrentBlindMulti())
	s.round = entity.NewPokerRound(
		s.runInfo.Deck,
		s.runInfo.DefaultHands,
		s.runInfo.DefaultDiscards,
		scoreAtLeast,
	)
	s.round.Deck.Shuffle()

	return nil
}

func (s *pokerService) DrawCard(num int) ([]entity.Trump, error) {
	cards := s.round.DrawCard(num)
	return cards, nil
}

func (s *pokerService) GetCurrentAnteAmount() int {
	return entity.AnteAmounts[s.runInfo.AnteIndex]
}

func (s *pokerService) NextRound() error {
	s.runInfo.NextRound()
	scoreAtLeast := int(float64(s.GetCurrentAnteAmount()) * s.GetCurrentBlindMulti())
	s.round = entity.NewPokerRound(
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
	if s.round.GetRoundStats().Discards > 0 {
		actions = append(actions, "Discard")
	}
	actions = append(actions, "Cancel")

	return actions
}

func (s *pokerService) SelectCards(cards []string) error {
	s.round.SetSelectCards(cards)

	return nil
}

func (s *pokerService) DiscardHand() error {
	s.round.Stats.Discards--

	return nil
}

func (s *pokerService) CancelHand() error {
	s.round.SelectedCards = nil

	return nil
}

func (s *pokerService) PlayHand() (entity.PokerHandStats, error) {
	s.round.Stats.Hands--

	round := s.round

	// get hand type and base chip and mult
	handType := round.PlayHand()
	chip, mult := s.GetChipAndMult(handType, 1)

	// get card rank and add to chip
	handsRankTotal := round.GetSelectCardsRankTotal()
	chip += handsRankTotal
	score := chip * mult
	round.Stats.TotalScore += score

	stats := entity.PokerHandStats{
		HandType: handType,
		Chip:     chip,
		Mult:     mult,
		Score:    score,
	}

	return stats, nil
}

func (s *pokerService) GetHandCardString() []string {
	return s.round.HandCardString()
}

func (s *pokerService) GetRemainCardString() []string {
	return s.round.RemainCardString()
}

func (s *pokerService) GetRoundStats() *entity.RoundStats {
	return s.round.GetRoundStats()
}

func (s *pokerService) SetAction(action string) {
	s.round.BeforeSelectAction = action
}

func (s *pokerService) IsRoundWin() bool {
	return s.round.IsWin()
}

func (s *pokerService) GetRounds() int {
	return s.runInfo.Rounds
}

func (s *pokerService) ShopOpen() []string {
	if s.runInfo.Rounds <= 1 {
		// return nil
	}

	shop := entity.NewShop()
	return shop.GetShopItems()
}

func (s *pokerService) AddShopItem(itemName string) error {
	jokerCards := entity.GetJokerCards()
	s.runInfo.AddJokerCard(jokerCards[itemName])

	return nil
}

func (s *pokerService) ShowJokers() {
	fmt.Printf("%v", s.runInfo.JokerCards)
}

// NewPokerServiceConfig returns a new PokerServiceConfig
func NewPokerServiceConfig() PokerServiceConfig {
	return PokerServiceConfig{}
}
