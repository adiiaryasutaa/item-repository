package util

import "os"

func Arguments() []string {
	return os.Args[1:]
}

func GetArgument(index int) string {
	arguments := Arguments()

	if len(arguments)-1 < index {
		return ""
	}

	return arguments[index]
}
