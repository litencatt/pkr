package entity

import (
	"crypto/rand"
	"math/big"
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
	for i := len(d) - 1; i > 0; i-- {
		// crypto/randを使用してセキュアな乱数を生成
		j, err := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
		if err != nil {
			// crypto/randが失敗した場合は何もしない（シャッフルなし）
			return
		}
		d[i], d[j.Int64()] = d[j.Int64()], d[i]
	}
}

func (d *Deck) Draw(n int) []Trump {
	hand := (*d)[:n]
	*d = (*d)[n:]
	return hand
}
