package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"item-repository/handler"
	"item-repository/util"
	"os"
)

func GetChoice() int {
	args := os.Args[1:]

	if len(args) != 0 {
		switch args[0] {
		case "add", "i":
			if len(args[1:]) < 3 {
				fmt.Println("Input kurang")
				return 0
			}

			return 1
		case "show", "d":
			return 2
		case "search", "s":
			if len(args) == 1 {
				fmt.Println("Input kurang")
				return 0
			}

			return 3
		case "about", "a":
			return 4
		default:
			return 0
		}
	}

	return menu()
}

func main() {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	for {
		choice := GetChoice()

		if choice == 0 {
			return
		}

		handler.Handle(uint8(choice))

		if util.ProgramRunDirectly() {
			break
		}
	}

	os.Exit(0)
}
