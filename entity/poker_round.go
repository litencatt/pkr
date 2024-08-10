package entity

import (
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
	HandCards          PokerHandCard
	RemainCards        []Trump
	SelectedCards      []Trump
	Hands              int
	Discards           int
	ScoreAtLeast       int
	BeforeSelectAction string
}

func (p *PokerRound) DrawCard(drawNum int) []Trump {
	if drawNum == 0 {
		return nil
	}

	// Remain cards
	p.HandCards.Trumps = nil
	p.HandCards.Trumps = append(p.HandCards.Trumps, p.RemainCards...)

	// Draw cards and append to hand
	drawCards := p.Deck.Draw(drawNum)
	p.HandCards.Trumps = append(p.HandCards.Trumps, drawCards...)

	p.HandCards.Sort()

	return drawCards
}

func (p *PokerRound) HandCardString() []string {
	var cards []string
	for _, card := range p.HandCards.Trumps {
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
		for _, t := range p.HandCards.Trumps {
			if string(t.Rank) == rank && string(t.Suit) == suit {
				selectCards = append(selectCards, t)
				break
			}
		}
	}
	p.SelectedCards = selectCards

	// Calc the RemainCards cards
	var RemainCardsCards []Trump
	for _, card := range p.HandCards.Trumps {
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
