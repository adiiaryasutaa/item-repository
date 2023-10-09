package util

import (
	"bufio"
	"fmt"
	"github.com/eiannone/keyboard"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func ReadInput() (string, error) {
	input, err := reader.ReadString('\n')
	input = strings.Replace(input, "\n", "", -1)
	input = strings.Replace(input, "\r", "", -1)

	return input, err
}

func Ask(question string) string {
	fmt.Print(question)

	input, _ := ReadInput()

	return input
}

func PressEnter() {
	for {
		_, key, err := keyboard.GetSingleKey()

		if err != nil {
			panic(err)
		}

		if key == keyboard.KeyEnter {
			break
		}
	}
}
