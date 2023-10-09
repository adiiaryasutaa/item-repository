package util

import "os"

// Arguments Get arguments after file name
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
