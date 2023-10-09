package main

import (
	"fmt"
	"item-repository/util"
	"item-repository/util/print"
	"strconv"
)

func menu() int {
	errorMessage := ""

	for {
		util.Clear()

		fmt.Println("ITEM REPOSITORY")
		fmt.Println()
		fmt.Println("1. Add Item")
		fmt.Println("2. Show Items")
		fmt.Println("3. Search Items")
		fmt.Println("4. About")
		fmt.Println()

		if errorMessage != "" {
			print.Error(errorMessage)

			errorMessage = ""
		}

		fmt.Print("> Choice menu [1 - 4]: ")

		var input string
		_, _ = fmt.Scanln(&input)

		if input == "" {
			continue
		}

		choice, err := strconv.Atoi(input)

		if err != nil {
			errorMessage = fmt.Sprintf("Invalid input [%s], please choice between 1 and 4\n", input)
			continue
		}

		if choice < 1 || choice > 4 {
			errorMessage = "Please choice between 1 and 4\n"
			continue
		}

		return choice
	}
}
