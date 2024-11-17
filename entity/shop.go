package entity

import (
	"strconv"
)

type ShopItemType string

const (
	Joker ShopItemType = "Joker"
)

type Shop struct {
	Items []ShopItem
}

type ShopItem struct {
	Name        string
	Description string
	Price       int
	Type        string
}

func NewShop() *Shop {
	var items []ShopItem

	joker := GetJokerCards()
	for name, jokerCard := range joker {
		items = append(items, ShopItem{
			Name:        name,
			Description: jokerCard.Effects.Description,
			Price:       100,
			Type:        string(Joker),
		})
	}

	return &Shop{Items: items}
}

func (s *Shop) GetShopItems() []string {
	var items []string
	for _, item := range s.Items {
		displayStr := item.Name + ": " + item.Description + " $" + strconv.Itoa(item.Price)
		items = append(items, displayStr)
	}

	return items
}
