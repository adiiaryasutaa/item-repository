package print

import (
	"fmt"
)

// Info Show info prefix
func Info(message string) {
	fmt.Println("| INFO |", message)
}

// Error Show error prefix
func Error(message string) {
	fmt.Println("| ERROR |", message)
}
