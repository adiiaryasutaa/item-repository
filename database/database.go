package database

import (
	"fmt"
	"os"
	"strings"
)

const filename = "./database/item.txt"

// Save For saving string to file
func Save(s string) int {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println("Error occurred when opening database file")
		os.Exit(1)
	}

	n, err := file.WriteString(s)

	if err != nil {
		fmt.Println("Error occurred when saving data")
		os.Exit(1)
	}

	if err := file.Close(); err != nil {
		fmt.Println(err)
	}

	return n
}

func Read() string {
	// Read the contents of the file into a slice of bytes.
	data, err := os.ReadFile(filename)

	if err != nil {
		fmt.Println(err)
		return ""
	}

	// Convert the slice of bytes to a string.
	return strings.Trim(string(data), "\r\n")
}
