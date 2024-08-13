package entity

import (
	"os"

	"gopkg.in/yaml.v2"
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
	Rarity      Rarity `yaml:"rarity"`
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Effect      func() `yaml:"-"`
}

func GetJokerCards() map[string]JokerCard {
	jokers, err := LoadJokerCardsFromJSON("./joker_common.yaml")
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

	var jokerList []JokerEffect
	err = yaml.Unmarshal(data, &jokerList)
	if err != nil {
		return nil, err
	}

	jokerCards := make(map[string]JokerCard)
	for _, card := range jokerList {
		jokerCards[card.Name] = JokerCard{Effects: &card}
	}

	return &jokerCards, nil
}
