package entity

import "strconv"

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
	return &Shop{
		Items: []ShopItem{
			{
				Name:        "Joker1",
				Description: "Add 1 multiples",
				Price:       100,
				Type:        string(Joker),
			},
		},
	}
}

func (s *Shop) GetShopItems() []string {
	var items []string
	for _, item := range s.Items {
		displayStr := item.Name + ": " + item.Description + " $" + strconv.Itoa(item.Price)
		items = append(items, displayStr)
	}

	return items
}
