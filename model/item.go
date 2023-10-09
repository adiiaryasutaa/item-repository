package model

import (
	"fmt"
	"item-repository/database"
	"item-repository/util"
	"strings"
)

type Item struct {
	Name     string
	Price    uint
	Quantity uint
}

func NewItem(name string, price uint, quantity uint) Item {
	return Item{Name: name, Price: price, Quantity: quantity}
}

func GetAllItem() []Item {
	data := database.Read()

	a := strings.Split(data, "\n")

	var items []Item

	for _, s := range a {
		s = strings.Trim(s, "\r\n")
		b := s[strings.LastIndex(s, "\"")+2:]
		g := strings.Split(b, ",")

		items = append(items, NewItem(s[1:strings.LastIndex(s, "\"")], util.ConvertToPrice(g[0]), util.ConvertToQuantity(g[1])))
	}

	return items
}

func FilterItems(query string) []Item {
	items := GetAllItem()
	var qualifiers []Item

	for _, item := range items {
		if strings.Contains(strings.ToLower(item.Name), strings.ToLower(query)) {
			qualifiers = append(qualifiers, item)
		}
	}

	return qualifiers
}

func PrintItemTable(items []Item) {
	fmt.Printf("%*s %*s %*s\n", 11, "NAME", 28, "PRICE (Rp)", 10, "QUANTITY")
	for i, item := range items {
		fmt.Printf("%*d | %*s | %*d | %*d\n", 4, i+1, -20, item.Name, 10, item.Price, len("Quantity"), item.Quantity)
	}
}
