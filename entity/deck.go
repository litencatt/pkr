package entity

import (
	"math/rand"
	"time"
)

type Deck []Trump

func NewDeck() Deck {
	deck := make(Deck, 0)
	suits := []Suit{Clubs, Diamonds, Hearts, Spades}
	ranks := []Rank{Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Ace}
	for _, suit := range suits {
		for _, rank := range ranks {
			deck = append(deck, Trump{Suit: suit, Rank: rank})
		}
	}
	return deck
}

func (d Deck) Len() int {
	return len(d)
}

func (d Deck) Shuffle() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
}

func (d *Deck) Draw(n int) []Trump {
	hand := (*d)[:n]
	*d = (*d)[n:]
	return hand
}
