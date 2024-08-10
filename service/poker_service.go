package service

import (
	"fmt"

	"github.com/litencatt/pkr/entity"
)

type PokerService interface {
	StartRound(int) error
	DrawCard(int) error
	PlayHand() (entity.PokerHandStats, error)
	DiscardHand() error
	NextRound() error
	NextAnte() error
	GetCurrentAnte() int
	GetCurrentBlind() float64
	GetNextDrawNum() int
	GetChipAndMult(entity.HandType, int) (int, int)
	SelectCards([]string) error
	GetHandCardString() []string
	GetRemainCardString() []string
	GetRoundStats() *entity.PokerRoundStats
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
	return s.runInfo.DefaultDeal
}

func (s *pokerService) GetChipAndMult(handType entity.HandType, level int) (int, int) {
	return s.runInfo.PokerHands.GetChipAndMult(handType, level)
}

func (s *pokerService) StartRound(ScoreAtLeast int) error {
	scoreAtLeast := int(float64(s.runInfo.Ante) * s.runInfo.Blind)
	s.runInfo.Round = &entity.PokerRound{
		Deck:         s.runInfo.Deck,
		TotalScore:   0,
		Hands:        s.runInfo.Hands,
		Discards:     s.runInfo.Discards,
		ScoreAtLeast: scoreAtLeast,
	}

	s.runInfo.Round.Deck.Shuffle()

	return nil
}

func (s *pokerService) DrawCard(nextDrawNum int) error {
	drawCards := s.runInfo.Round.DrawCard(nextDrawNum)
	if s.config.DebugMode {
		fmt.Println("[Draw", nextDrawNum, "cards]")
		for _, card := range drawCards {
			fmt.Println(card.String())
		}
		fmt.Println()
	}

	return nil
}

func (s *pokerService) GetCurrentAnte() int {
	return s.runInfo.Ante
}

func (s *pokerService) GetCurrentBlind() float64 {
	return s.runInfo.Blind
}

func (s *pokerService) SelectCards(cards []string) error {
	s.runInfo.Round.SetSelectCards(cards)
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

func (s *pokerService) DiscardHand() error {
	s.runInfo.Discards--
	return nil
}

func (s *pokerService) GetHandCardString() []string {
	return s.runInfo.Round.HandCardString()
}

func (s *pokerService) GetRemainCardString() []string {
	return s.runInfo.Round.RemainCardString()
}

func (s *pokerService) NextRound() error {
	return nil
}

func (s *pokerService) NextAnte() error {
	return nil
}

func (s *pokerService) GetRoundStats() *entity.PokerRoundStats {
	return s.runInfo.Round.GetRoundStats()
}

// NewPokerServiceConfig returns a new PokerServiceConfig
func NewPokerServiceConfig() PokerServiceConfig {
	return PokerServiceConfig{}
}