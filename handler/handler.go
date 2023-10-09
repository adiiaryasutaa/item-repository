package handler

import (
	"fmt"
	"item-repository/model"
	"item-repository/util"
	"item-repository/util/print"
)

// Handle Handling choices
func Handle(choice uint8) {
	if !util.ProgramRunDirectly() {
		util.Clear()
	}

	switch choice {
	case 1: // Add
		m := model.NewItem(CollectItemData())

		AddItem(m)

		fmt.Println()
		print.Info(fmt.Sprintf("Item [Name: %s, Price: %d, Quantity: %d] Added", m.Name, m.Price, m.Quantity))

		if util.ProgramRunDirectly() {
			return
		}

	case 2: // Show
		items := model.GetAllItem()

		fmt.Printf("\n%*s\n\n", -3, "Item List")

		model.PrintItemTable(items)

		if util.ProgramRunDirectly() {
			return
		}

	case 3: // Search
		var query string

		fmt.Println("Search Item")
		fmt.Println()

		if util.ProgramRunDirectly() {
			query = util.GetArgument(1)
		} else {
			fmt.Print("> Enter Item Name: ")
			input, err := util.ReadInput()

			if err != nil {
				fmt.Println(err)
				return
			}

			query = input
		}

		items := model.FilterItems(query)

		fmt.Println()
		if len(items) == 0 {
			print.Info("Item Not Found")
			break
		} else {
			model.PrintItemTable(items)
		}

		if util.ProgramRunDirectly() {
			return
		}

	case 4: // About
		ShowAbout()

		if util.ProgramRunDirectly() {
			return
		}

	default:
		fmt.Println("Invalid Action")
	}

	fmt.Println()
	fmt.Println("(Press Enter to back)")
	util.PressEnter()
}
