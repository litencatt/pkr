package entity

import (
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := NewDeck()

	if deck == nil {
		t.Fatal("NewDeck() returned nil")
	}

	if len(deck) != 52 {
		t.Errorf("NewDeck() created %d cards, want 52", len(deck))
	}

	// Check for uniqueness of cards
	seen := make(map[Trump]bool)
	for _, card := range deck {
		if seen[card] {
			t.Errorf("Duplicate card found: %s", card.String())
		}
		seen[card] = true
	}

	// Check that all suits and ranks are present
	allSuits := []Suit{Spades, Hearts, Diamonds, Clubs}
	allRanks := []Rank{Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Ace}
	for _, suit := range allSuits {
		for _, rank := range allRanks {
			card := Trump{Suit: suit, Rank: rank}
			if !seen[card] {
				t.Errorf("Missing card: %s", card.String())
			}
		}
	}
}

func TestDeckShuffle(t *testing.T) {
	deck1 := NewDeck()

	// Save original order
	originalOrder := make([]Trump, 52)
	copy(originalOrder, deck1)

	// Shuffle deck1
	deck1.Shuffle()

	// Check that all cards are still present
	if len(deck1) != 52 {
		t.Errorf("After shuffle, deck has %d cards, want 52", len(deck1))
	}

	// Check that the order has changed (with very high probability)
	same := true
	for i := 0; i < 52; i++ {
		if originalOrder[i] != deck1[i] {
			same = false
			break
		}
	}

	if same {
		t.Log("Warning: Deck order unchanged after shuffle (very unlikely)")
	}

	// Ensure all cards are still unique
	seen := make(map[Trump]bool)
	for _, card := range deck1 {
		if seen[card] {
			t.Errorf("Duplicate card after shuffle: %s", card.String())
		}
		seen[card] = true
	}
}

func TestDeckDraw(t *testing.T) {
	deck := NewDeck()
	deck.Shuffle()

	// Test drawing cards
	drawn := deck.Draw(5)

	if len(drawn) != 5 {
		t.Errorf("Draw(5) returned %d cards, want 5", len(drawn))
	}

	if len(deck) != 47 {
		t.Errorf("After drawing 5 cards, deck has %d cards, want 47", len(deck))
	}

	// Test drawing more cards
	drawn2 := deck.Draw(10)

	if len(drawn2) != 10 {
		t.Errorf("Draw(10) returned %d cards, want 10", len(drawn2))
	}

	if len(deck) != 37 {
		t.Errorf("After drawing 15 cards total, deck has %d cards, want 37", len(deck))
	}

	// Test drawing more cards
	drawn3 := deck.Draw(30)

	if len(drawn3) != 30 {
		t.Errorf("Draw(30) returned %d cards, want 30", len(drawn3))
	}

	if len(deck) != 7 {
		t.Errorf("After drawing 45 cards total, deck has %d cards, want 7", len(deck))
	}
}

func TestDeckLen(t *testing.T) {
	deck := NewDeck()

	if deck.Len() != 52 {
		t.Errorf("NewDeck().Len() = %d, want 52", deck.Len())
	}

	// Draw some cards
	deck.Draw(5)

	if deck.Len() != 47 {
		t.Errorf("After drawing 5 cards, deck.Len() = %d, want 47", deck.Len())
	}
}
