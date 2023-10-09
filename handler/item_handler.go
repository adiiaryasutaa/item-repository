package handler

import (
	"fmt"
	"item-repository/database"
	"item-repository/model"
	"item-repository/util"
)

func AddItem(item model.Item) {
	database.Save(fmt.Sprintf("\"%s\",%d,%d\n", item.Name, item.Price, item.Quantity))
}

func CollectItemData() (string, uint, uint) {
	var (
		name     string
		price    uint
		quantity uint
	)

	if util.ProgramRunDirectly() {
		name = util.GetArgument(1)
		price = util.ConvertToPrice(util.GetArgument(2))
		quantity = util.ConvertToQuantity(util.GetArgument(3))
	} else {
		fmt.Println("ADD ITEM")
		fmt.Println()

		name = util.Ask("> Item Name: ")
		price = util.ConvertToPrice(util.Ask("> Item Price: "))
		quantity = util.ConvertToQuantity(util.Ask("> Item Quantity: "))
	}

	return name, price, quantity
}
