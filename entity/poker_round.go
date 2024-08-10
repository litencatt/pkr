package entity

import (
	"sort"
	"strings"
)

type PokerRoundStats struct {
	Hands        int
	Discards     int
	TotalScore   int
	ScoreAtLeast int
}

type PokerRound struct {
	Deck               Deck
	TotalScore         int
	HandCards          []Trump
	RemainCards        []Trump
	SelectedCards      []Trump
	Hands              int
	Discards           int
	ScoreAtLeast       int
	BeforeSelectAction string
}

func NewPokerRound(deck Deck, hands, discards, scoreAtLeast int) *PokerRound {
	return &PokerRound{
		Deck:          deck,
		TotalScore:    0,
		HandCards:     nil,
		RemainCards:   nil,
		SelectedCards: nil,
		Hands:         hands,
		Discards:      discards,
		ScoreAtLeast:  scoreAtLeast,
	}
}

func SortTrumps(trumps []Trump) {
	sort.Slice(trumps, func(i, j int) bool {
		return trumps[i].GetSortOrder() < trumps[j].GetSortOrder()
	})
}

func (p *PokerRound) DrawCard(drawNum int) []Trump {
	if drawNum == 0 {
		return nil
	}

	// Remain cards
	p.HandCards = nil
	p.HandCards = append(p.HandCards, p.RemainCards...)

	// Draw cards and append to hand
	drawCards := p.Deck.Draw(drawNum)
	p.HandCards = append(p.HandCards, drawCards...)

	// Sort hand cards
	SortTrumps(p.HandCards)

	return drawCards
}

func (p *PokerRound) HandCardString() []string {
	var cards []string
	for _, card := range p.HandCards {
		cards = append(cards, card.String())
	}
	return cards
}

func (p *PokerRound) RemainCardString() []string {
	var cards []string
	for _, card := range p.RemainCards {
		cards = append(cards, card.String())
	}
	return cards
}

func (p *PokerRound) SetSelectCards(cards []string) {
	// Convert select cards to Trump entity
	var selectCards []Trump
	for _, card := range cards {
		// extract rank and suit from card string
		rank := strings.Split(card, " of ")[0]
		suit := strings.Split(card, " of ")[1]
		// Find the card from hand
		for _, t := range p.HandCards {
			if string(t.Rank) == rank && string(t.Suit) == suit {
				selectCards = append(selectCards, t)
				break
			}
		}
	}
	p.SelectedCards = selectCards

	// Calc the RemainCards cards
	var RemainCardsCards []Trump
	for _, card := range p.HandCards {
		if !Contains(selectCards, card) {
			RemainCardsCards = append(RemainCardsCards, card)
		}
	}
	p.RemainCards = RemainCardsCards
}

func (p *PokerRound) PlayHand() HandType {
	return EvaluateHand(p.SelectedCards)
}

func (p *PokerRound) GetSelectCardsRankTotal() int {
	total := 0
	for _, card := range p.SelectedCards {
		total += card.GetRankNumber()
	}
	return total
}

func (p *PokerRound) GetRoundStats() *PokerRoundStats {
	return &PokerRoundStats{
		Hands:        p.Hands,
		Discards:     p.Discards,
		TotalScore:   p.TotalScore,
		ScoreAtLeast: p.ScoreAtLeast,
	}
}

func (p *PokerRound) IsWin() bool {
	return p.TotalScore >= p.ScoreAtLeast
}
