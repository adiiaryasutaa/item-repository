package util

import (
	"os"
	"os/exec"
	"runtime"
)

func ProgramRunDirectly() bool {
	return len(os.Args[1:]) > 0
}

func Clear() {
	var cmd *exec.Cmd

	// Clear the terminal on Linux and macOS.
	if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		cmd = exec.Command("clear")
	}

	// Clear the terminal on Windows.
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}
