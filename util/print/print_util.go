package print

import (
	"fmt"
)

func Info(message string) {
	fmt.Println("| INFO |", message)
}

func Error(message string) {
	fmt.Println("| ERROR |", message)
}
