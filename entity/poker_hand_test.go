package entity

import (
	"testing"
)

func TestNewPokerHands(t *testing.T) {
	hands := NewPokerHands()

	if hands == nil {
		t.Fatal("NewPokerHands() returned nil")
	}

	// Check that all hand types are present
	expectedHands := []HandType{
		HighCard, OnePair, TwoPair, ThreeOfAKind,
		Straight, Flush, FullHouse, FourOfAKind,
		StraightFlush, RoyalFlush,
	}

	if len(hands.PokerHands) != len(expectedHands) {
		t.Errorf("Expected %d hand types, got %d", len(expectedHands), len(hands.PokerHands))
	}

	handTypeMap := make(map[HandType]bool)
	for _, hand := range hands.PokerHands {
		handTypeMap[hand.HandType] = true
	}

	for _, expected := range expectedHands {
		if !handTypeMap[expected] {
			t.Errorf("Missing hand type: %s", expected)
		}
	}
}

func TestGetChipAndMult(t *testing.T) {
	hands := NewPokerHands()

	tests := []struct {
		name     string
		handType HandType
		level    int
		wantChip int
		wantMult int
	}{
		{
			name:     "High Card Level 1",
			handType: HighCard,
			level:    1,
			wantChip: 5,
			wantMult: 1,
		},
		{
			name:     "One Pair Level 1",
			handType: OnePair,
			level:    1,
			wantChip: 10,
			wantMult: 2,
		},
		{
			name:     "Two Pair Level 1",
			handType: TwoPair,
			level:    1,
			wantChip: 20,
			wantMult: 2,
		},
		{
			name:     "Flush Level 1",
			handType: Flush,
			level:    1,
			wantChip: 35,
			wantMult: 4,
		},
		{
			name:     "Royal Flush Level 1",
			handType: RoyalFlush,
			level:    1,
			wantChip: 100,
			wantMult: 8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			chip, mult := hands.GetChipAndMult(tt.handType, tt.level)
			if chip != tt.wantChip {
				t.Errorf("GetChipAndMult() chip = %d, want %d", chip, tt.wantChip)
			}
			if mult != tt.wantMult {
				t.Errorf("GetChipAndMult() mult = %d, want %d", mult, tt.wantMult)
			}
		})
	}
}

func TestGetChipAndMultInvalidHandType(t *testing.T) {
	hands := NewPokerHands()

	// Test with invalid hand type
	chip, mult := hands.GetChipAndMult("InvalidHand", 1)

	// Should return 0 for both when hand type doesn't exist
	if chip != 0 || mult != 0 {
		t.Errorf("GetChipAndMult() with invalid hand returned chip=%d, mult=%d, want 0, 0", chip, mult)
	}
}

func TestGetChipAndMultInvalidLevel(t *testing.T) {
	hands := NewPokerHands()

	// Test with invalid level (level 0 or negative)
	chip, mult := hands.GetChipAndMult(OnePair, 0)
	if chip != 0 || mult != 0 {
		t.Errorf("GetChipAndMult() with level 0 returned chip=%d, mult=%d, want 0, 0", chip, mult)
	}

	// Test with level higher than available
	chip, mult = hands.GetChipAndMult(OnePair, 999)
	if chip != 0 || mult != 0 {
		t.Errorf("GetChipAndMult() with level 999 returned chip=%d, mult=%d, want 0, 0", chip, mult)
	}
}

func TestPokerHandStatsCalculation(t *testing.T) {
	tests := []struct {
		name      string
		handStats PokerHandStats
		wantScore int
	}{
		{
			name: "High Card Score",
			handStats: PokerHandStats{
				HandType: HighCard,
				Level:    1,
				Chip:     5,
				Mult:     1,
			},
			wantScore: 5,
		},
		{
			name: "One Pair Score",
			handStats: PokerHandStats{
				HandType: OnePair,
				Level:    1,
				Chip:     10,
				Mult:     2,
			},
			wantScore: 20,
		},
		{
			name: "Flush Score",
			handStats: PokerHandStats{
				HandType: Flush,
				Level:    1,
				Chip:     35,
				Mult:     4,
			},
			wantScore: 140,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Calculate score
			tt.handStats.Score = tt.handStats.Chip * tt.handStats.Mult

			if tt.handStats.Score != tt.wantScore {
				t.Errorf("PokerHandStats.Score = %d, want %d", tt.handStats.Score, tt.wantScore)
			}
		})
	}
}
