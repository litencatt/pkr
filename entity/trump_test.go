package entity

import (
	"testing"
)

func TestTrumpString(t *testing.T) {
	tests := []struct {
		name string
		card Trump
		want string
	}{
		{
			name: "Ace of Spades",
			card: Trump{Suit: Spades, Rank: Ace},
			want: "A of Spades",
		},
		{
			name: "King of Hearts",
			card: Trump{Suit: Hearts, Rank: King},
			want: "K of Hearts",
		},
		{
			name: "Queen of Diamonds",
			card: Trump{Suit: Diamonds, Rank: Queen},
			want: "Q of Diamonds",
		},
		{
			name: "Jack of Clubs",
			card: Trump{Suit: Clubs, Rank: Jack},
			want: "J of Clubs",
		},
		{
			name: "10 of Spades",
			card: Trump{Suit: Spades, Rank: Ten},
			want: "T of Spades",
		},
		{
			name: "2 of Hearts",
			card: Trump{Suit: Hearts, Rank: Two},
			want: "2 of Hearts",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.card.String(); got != tt.want {
				t.Errorf("Trump.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrumpEquality(t *testing.T) {
	card1 := Trump{Suit: Spades, Rank: Ace}
	card2 := Trump{Suit: Spades, Rank: Ace}
	card3 := Trump{Suit: Hearts, Rank: Ace}

	if card1 != card2 {
		t.Error("Same cards should be equal")
	}

	if card1 == card3 {
		t.Error("Different cards should not be equal")
	}
}

func TestTrumpGetRankNumber(t *testing.T) {
	tests := []struct {
		name string
		card Trump
		want int
	}{
		{"Ace", Trump{Rank: Ace}, 14},
		{"King", Trump{Rank: King}, 13},
		{"Queen", Trump{Rank: Queen}, 12},
		{"Jack", Trump{Rank: Jack}, 11},
		{"Ten", Trump{Rank: Ten}, 10},
		{"Nine", Trump{Rank: Nine}, 9},
		{"Two", Trump{Rank: Two}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.card.GetRankNumber(); got != tt.want {
				t.Errorf("GetRankNumber() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestTrumpGetSortOrder(t *testing.T) {
	tests := []struct {
		name string
		card Trump
		want int
	}{
		{"Ace", Trump{Rank: Ace}, 14},
		{"King", Trump{Rank: King}, 13},
		{"Queen", Trump{Rank: Queen}, 12},
		{"Jack", Trump{Rank: Jack}, 11},
		{"Ten", Trump{Rank: Ten}, 10},
		{"Nine", Trump{Rank: Nine}, 9},
		{"Two", Trump{Rank: Two}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.card.GetSortOrder(); got != tt.want {
				t.Errorf("GetSortOrder() = %d, want %d", got, tt.want)
			}
		})
	}
}
