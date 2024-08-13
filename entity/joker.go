package entity

import (
	"encoding/json"
	"os"
)

type Rarity string

const (
	Common   Rarity = "Common"
	Uncommon Rarity = "Uncommon"
	Rare     Rarity = "Rare"
	Legend   Rarity = "Legend"
)

type JokerCard struct {
	Effects *JokerEffect
}

type JokerEffect struct {
	Rarity      Rarity `json:"rarity"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Effect      func() `json:"-"`
}

func GetJokerCards() map[string]JokerCard {
	jokers, err := LoadJokerCardsFromJSON("./joker_common.json")
	if err != nil {
		return nil
	}

	return *jokers
}

func LoadJokerCardsFromJSON(filename string) (*map[string]JokerCard, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var jokerData map[string]JokerEffect
	err = json.Unmarshal(data, &jokerData)
	if err != nil {
		return nil, err
	}

	jokerCards := make(map[string]JokerCard)
	for key, effect := range jokerData {
		jokerCards[key] = JokerCard{Effects: &effect}
	}

	return &jokerCards, nil
}
