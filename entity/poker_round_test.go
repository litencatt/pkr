package entity

import (
	"testing"
)

func TestNewPokerRound(t *testing.T) {
	deck := NewDeck()
	hands := 4
	discards := 3
	scoreAtLeast := 300

	round := NewPokerRound(deck, hands, discards, scoreAtLeast)

	if round == nil {
		t.Fatal("NewPokerRound() returned nil")
	}

	stats := round.GetRoundStats()
	if stats.Hands != hands {
		t.Errorf("NewPokerRound() Hands = %d, want %d", stats.Hands, hands)
	}

	if stats.Discards != discards {
		t.Errorf("NewPokerRound() Discards = %d, want %d", stats.Discards, discards)
	}

	if stats.ScoreAtLeast != scoreAtLeast {
		t.Errorf("NewPokerRound() ScoreAtLeast = %d, want %d", stats.ScoreAtLeast, scoreAtLeast)
	}

	if stats.TotalScore != 0 {
		t.Errorf("NewPokerRound() TotalScore = %d, want 0", stats.TotalScore)
	}

	if len(round.HandCards) != 0 {
		t.Errorf("NewPokerRound() HandCards has %d cards, want 0", len(round.HandCards))
	}

	if len(round.SelectedCards) != 0 {
		t.Errorf("NewPokerRound() SelectedCards has %d cards, want 0", len(round.SelectedCards))
	}
}

func TestPokerRoundDrawCard(t *testing.T) {
	deck := NewDeck()
	deck.Shuffle()
	round := NewPokerRound(deck, 4, 3, 300)

	// Draw initial cards
	drawn := round.DrawCard(8)

	if len(drawn) != 8 {
		t.Errorf("DrawCard(8) returned %d cards, want 8", len(drawn))
	}

	if len(round.HandCards) != 8 {
		t.Errorf("After drawing 8 cards, HandCards has %d cards, want 8", len(round.HandCards))
	}

	// Save some cards as remain cards (simulate selecting cards)
	round.RemainCards = round.HandCards[:3]
	round.HandCards = nil

	// Draw more cards (should add to remain cards)
	drawn2 := round.DrawCard(5)

	if len(drawn2) != 5 {
		t.Errorf("DrawCard(5) returned %d cards, want 5", len(drawn2))
	}

	if len(round.HandCards) != 8 {
		t.Errorf("After drawing with 3 remain cards, HandCards has %d cards, want 8", len(round.HandCards))
	}
}

func TestPokerRoundHandCardString(t *testing.T) {
	deck := NewDeck()
	round := NewPokerRound(deck, 4, 3, 300)

	// Add specific cards to hand
	round.HandCards = []Trump{
		{Suit: Spades, Rank: Ace},
		{Suit: Hearts, Rank: King},
		{Suit: Diamonds, Rank: Queen},
	}

	cardStrings := round.HandCardString()

	if len(cardStrings) != 3 {
		t.Errorf("HandCardString() returned %d strings, want 3", len(cardStrings))
	}

	expectedStrings := []string{"A of Spades", "K of Hearts", "Q of Diamonds"}
	for i, expected := range expectedStrings {
		if cardStrings[i] != expected {
			t.Errorf("HandCardString()[%d] = %s, want %s", i, cardStrings[i], expected)
		}
	}
}

func TestPokerRoundSetSelectCards(t *testing.T) {
	deck := NewDeck()
	round := NewPokerRound(deck, 4, 3, 300)

	// Add cards to hand
	round.HandCards = []Trump{
		{Suit: Spades, Rank: Ace},
		{Suit: Hearts, Rank: King},
		{Suit: Diamonds, Rank: Queen},
		{Suit: Clubs, Rank: Jack},
		{Suit: Spades, Rank: Ten},
	}

	// Select some cards
	selectedStrings := []string{"A of Spades", "K of Hearts", "T of Spades"}
	round.SetSelectCards(selectedStrings)

	if len(round.SelectedCards) != 3 {
		t.Errorf("After selecting 3 cards, SelectedCards has %d cards, want 3", len(round.SelectedCards))
	}

	// Verify selected cards are correct
	expectedSelected := []Trump{
		{Suit: Spades, Rank: Ace},
		{Suit: Hearts, Rank: King},
		{Suit: Spades, Rank: Ten},
	}

	for i, expected := range expectedSelected {
		if round.SelectedCards[i] != expected {
			t.Errorf("SelectedCards[%d] = %v, want %v", i, round.SelectedCards[i], expected)
		}
	}

	// Check that remaining cards are correct
	if len(round.RemainCards) != 2 {
		t.Errorf("RemainCards has %d cards, want 2", len(round.RemainCards))
	}
}

func TestPokerRoundGetSelectCardsRankTotal(t *testing.T) {
	deck := NewDeck()
	round := NewPokerRound(deck, 4, 3, 300)

	round.SelectedCards = []Trump{
		{Suit: Spades, Rank: Ace},     // 14
		{Suit: Hearts, Rank: King},    // 13
		{Suit: Diamonds, Rank: Queen}, // 12
		{Suit: Clubs, Rank: Jack},     // 11
		{Suit: Spades, Rank: Nine},    // 9
	}

	total := round.GetSelectCardsRankTotal()
	expectedTotal := 14 + 13 + 12 + 11 + 9 // Ace + King + Queen + Jack + Nine
	if total != expectedTotal {
		t.Errorf("GetSelectCardsRankTotal() = %d, want %d", total, expectedTotal)
	}
}

func TestPokerRoundIsWin(t *testing.T) {
	deck := NewDeck()
	round := NewPokerRound(deck, 4, 3, 300)

	// Initially should not be win
	if round.IsWin() {
		t.Error("IsWin() should be false initially")
	}

	// Set score to meet requirement
	round.Stats.TotalScore = 300
	if !round.IsWin() {
		t.Error("IsWin() should be true when TotalScore >= ScoreAtLeast")
	}

	// Set score above requirement
	round.Stats.TotalScore = 350
	if !round.IsWin() {
		t.Error("IsWin() should be true when TotalScore > ScoreAtLeast")
	}

	// Set score below requirement
	round.Stats.TotalScore = 299
	if round.IsWin() {
		t.Error("IsWin() should be false when TotalScore < ScoreAtLeast")
	}
}

func TestRoundStats(t *testing.T) {
	deck := NewDeck()
	round := NewPokerRound(deck, 4, 3, 300)

	round.Stats.TotalScore = 150
	round.Stats.Hands = 2
	round.Stats.Discards = 1

	stats := round.GetRoundStats()

	if stats.ScoreAtLeast != 300 {
		t.Errorf("RoundStats.ScoreAtLeast = %d, want 300", stats.ScoreAtLeast)
	}

	if stats.TotalScore != 150 {
		t.Errorf("RoundStats.TotalScore = %d, want 150", stats.TotalScore)
	}

	if stats.Hands != 2 {
		t.Errorf("RoundStats.Hands = %d, want 2", stats.Hands)
	}

	if stats.Discards != 1 {
		t.Errorf("RoundStats.Discards = %d, want 1", stats.Discards)
	}
}
