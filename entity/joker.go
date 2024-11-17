package entity

import (
	"fmt"
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
	Rarity      Rarity      `yaml:"rarity"`
	Name        string      `yaml:"name"`
	Description string      `yaml:"description"`
	Price       int         `yaml:"price"`
	Condition   []Condition `yaml:"apply_conditions"`
	Effect      []Effect    `yaml:"effects"`
}

type Condition struct {
	Rank      Rank      `yaml:"rank"`
	Suit      Suit      `yaml:"suit"`
	PokerHand PokerHand `yaml:"poker_hand"`
}

type Effect struct {
	Chips     *ChipsEffect     `yaml:"chips,omitempty"`
	Multiples *MultiplesEffect `yaml:"multiples,omitempty"`
}

type ChipsEffect struct {
	Add int `yaml:"add"`
}

type MultiplesEffect struct {
	Add      int `yaml:"add,omitempty"`
	Multiply int `yaml:"multiply,omitempty"`
}

func (e *JokerEffect) IsApplicable(cards []Trump) bool {
	return true
}

func ApplyEffect(effect *JokerEffect, chips *int, mult *int) {
	if effect == nil {
		return
	}

	for _, eff := range effect.Effect {
		if eff.Chips != nil {
			*chips += eff.Chips.Add
		} else {
			if eff.Multiples.Add != 0 {
				*mult += eff.Multiples.Add
			} else {
				*mult *= eff.Multiples.Multiply
			}
		}
	}
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
